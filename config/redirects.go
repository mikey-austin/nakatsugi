package config

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/mikey-austin/nakatsugi/parser"
)

func SetRstmt(ctx parser.IRstmtContext, redirect *Redirect) error {
	if ctx.DISABLE() != nil {
		redirect.Disabled = true
	}

	if ctx.STICKY_ADDRESS() != nil {
		redirect.StickyAddress = true
	}

	if ctx.Sesstimeout() != nil {
		parsedTimeout, err := strconv.Atoi(ctx.Sesstimeout().NUMBER().GetText())
		if err != nil {
			return err
		}
		redirect.SessionTimeout = time.Duration(parsedTimeout) * time.Second
	}

	if ctx.Listen() != nil {
		err := SetListen(ctx.Listen(), redirect)
		if err != nil {
			return err
		}
	}

	return nil
}

func SetListen(ctx parser.IListenContext, redirect *Redirect) error {
	var listen Listen

	address := net.ParseIP(ctx.NAME().GetText())
	if address == nil {
		return fmt.Errorf("invalid listen address: %v", ctx.NAME().GetText())
	}
	listen.Address = address

	listen.Proto = ProtoTcp
	if ctx.Proto() != nil {
		switch ctx.Proto().NAME().GetText() {
		case "tcp":
			listen.Proto = ProtoTcp
		case "udp":
			listen.Proto = ProtoUdp
		default:
			return fmt.Errorf("unknown proto: %v", ctx.Proto().GetText())
		}
	}

	listen.Interface = ""
	if ctx.Interface_() != nil {
		iname := ctx.Interface_().NAME().GetText()
		if !IsValidName(iname) {
			return fmt.Errorf("invalid interface name: %v", iname)
		}
		listen.Interface = iname
	}

	listen.MinPort, listen.MaxPort = 0, 0
	if lport := ctx.Lport(); lport != nil {
		if r := lport.Portrange(); r != nil {
			name := r.NAME().GetText()
			parts := strings.Split(name, ":")
			if len(parts) != 2 {
				return fmt.Errorf("invalid port range: %v", name)
			}

			min, err := strconv.ParseUint(parts[0], 10, 16)
			if err != nil {
				return err
			}

			max, err := strconv.ParseUint(parts[1], 10, 16)
			if err != nil {
				return err
			}

			if min > max {
				return fmt.Errorf("min port (%v) not <= max port (%v)", min, max)
			}

			if !IsValidPort(min) || !IsValidPort(max) {
				return fmt.Errorf("invalid port range: %v:%v", min, max)
			}

			listen.MinPort = uint16(min)
			listen.MaxPort = uint16(max)
		} else if p := lport.Port(); p != nil {
			port, err := strconv.ParseUint(p.NUMBER().GetText(), 10, 16)
			if err != nil {
				return err
			}

			if !IsValidPort(port) {
				return fmt.Errorf("invalid port: %v", port)
			}

			listen.MinPort = uint16(port)
			listen.MaxPort = listen.MinPort
		}
	}

	redirect.Listen = append(redirect.Listen, &listen)

	return nil
}
