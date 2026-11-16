package config

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mikey-austin/nakatsugi/parser"
)

func NewListen(ctx parser.IListenContext) (*Listen, error) {
	var listen Listen

	address := net.ParseIP(ctx.NAME().GetText())
	if address == nil {
		return nil, fmt.Errorf("invalid listen address: %v", ctx.NAME().GetText())
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
			return nil, fmt.Errorf("unknown proto: %v", ctx.Proto().GetText())
		}
	}

	listen.Interface = ""
	if ctx.Interface_() != nil {
		iname := ctx.Interface_().NAME().GetText()
		if !IsValidName(iname) {
			return nil, fmt.Errorf("invalid interface name: %v", iname)
		}
		listen.Interface = iname
	}

	listen.MinPort, listen.MaxPort = 0, 0
	if lport := ctx.Lport(); lport != nil {
		if r := lport.Portrange(); r != nil {
			name := r.NAME().GetText()
			parts := strings.Split(name, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid port range: %v", name)
			}

			min, err := strconv.ParseUint(parts[0], 10, 16)
			if err != nil {
				return nil, err
			}

			max, err := strconv.ParseUint(parts[1], 10, 16)
			if err != nil {
				return nil, err
			}

			if min > max {
				return nil, fmt.Errorf("min port (%v) not <= max port (%v)", min, max)
			}

			if !IsValidPort(min) || !IsValidPort(max) {
				return nil, fmt.Errorf("invalid port range: %v:%v", min, max)
			}

			listen.MinPort = uint16(min)
			listen.MaxPort = uint16(max)
		} else if p := lport.Port(); p != nil {
			port, err := strconv.ParseUint(p.NUMBER().GetText(), 10, 16)
			if err != nil {
				return nil, err
			}

			if !IsValidPort(port) {
				return nil, fmt.Errorf("invalid port: %v", port)
			}

			listen.MinPort = uint16(port)
			listen.MaxPort = listen.MinPort
		}
	}

	return &listen, nil
}

func NewForward(ctx parser.IForwardContext, tableResolver func(ctx parser.ITnameContext) (*Table, error), defaultPort uint16) (*Forward, error) {
	destination, error := NewDestination(ctx.Destination(), tableResolver, defaultPort)
	if error != nil {
		return nil, error
	}
	forward := Forward{Destination: destination}
	return &forward, nil
}

func NewRoute(ctx parser.IRouteContext, tableResolver func(ctx parser.ITnameContext) (*Table, error), defaultPort uint16) (*Route, error) {
	destination, error := NewDestination(ctx.Destination(), tableResolver, defaultPort)
	if error != nil {
		return nil, error
	}
	route := Route{Destination: destination}
	return &route, nil
}

func NewDestination(ctx parser.IDestinationContext, tableResolver func(ctx parser.ITnameContext) (*Table, error), defaultPort uint16) (*Destination, error) {
	table, err := tableResolver(ctx.Tname())
	if err != nil {
		return nil, err
	}

	destination := Destination{ToTable: table}
	destination.Port = defaultPort
	if ctx.Port() != nil {
		p, err := strconv.ParseUint(ctx.Port().NUMBER().GetText(), 10, 16)
		if err != nil {
			return nil, err
		}
		destination.Port = uint16(p)
	}

	for _, dopt := range ctx.AllDopts() {
		switch {
		case dopt.Check() != nil:
			check, err := NewCheck(dopt.Check())
			if err != nil {
				return nil, err
			}
			destination.Check = check
			break
		}
	}

	return &destination, nil
}

func unquote(qname antlr.TerminalNode) string {
	return strings.Trim(qname.GetText(), "\"")
}

func host(ctx parser.IChostContext) (string, error) {
	if ctx == nil {
		return "", nil
	}

	hostname := ctx.NAME().GetText()
	if !IsValidHostname(hostname) {
		return "", fmt.Errorf("invalid hostname: %v", hostname)
	}
	return hostname, nil
}

func code(number antlr.TerminalNode) (uint, error) {
	code, err := strconv.ParseUint(number.GetText(), 10, 16)
	if err != nil {
		return uint(code), err
	}
	return uint(code), nil
}

func NewCheck(ctx parser.ICheckContext) (*Check, error) {
	var check Check

	// TODO: add general override options and mode options

	ct := ctx.Ctype()
	switch {
	case ct.Chttpcode() != nil:
		c := ct.Chttpcode()
		check.Type = CheckTypeHttpCode
		host, err := host(c.Chost())
		if err != nil {
			return nil, err
		}
		code, err := code(c.NUMBER())
		if err != nil {
			return nil, err
		}
		base := CheckHttpBase{Path: unquote(c.QNAME()), Hostname: host, Https: false}
		check.CheckHttpCode = &CheckHttpCode{CheckHttpBase: base, Code: code}
		break
	case ct.Chttpdigest() != nil:
		c := ct.Chttpdigest()
		check.Type = CheckTypeHttpDigest
		host, err := host(c.Chost())
		if err != nil {
			return nil, err
		}
		base := CheckHttpBase{Path: unquote(c.QNAME()), Hostname: host, Https: false}
		check.CheckHttpDigest = &CheckHttpDigest{CheckHttpBase: base, Digest: unquote(c.Cdigest().QNAME())}
		break
	case ct.Chttpscode() != nil:
		c := ct.Chttpscode()
		check.Type = CheckTypeHttpCode
		host, err := host(c.Chost())
		if err != nil {
			return nil, err
		}
		code, err := code(c.NUMBER())
		if err != nil {
			return nil, err
		}
		base := CheckHttpBase{Path: unquote(c.QNAME()), Hostname: host, Https: true}
		check.CheckHttpCode = &CheckHttpCode{CheckHttpBase: base, Code: code}
		break
	case ct.Chttpsdigest() != nil:
		c := ct.Chttpsdigest()
		check.Type = CheckTypeHttpDigest
		host, err := host(c.Chost())
		if err != nil {
			return nil, err
		}
		base := CheckHttpBase{Path: unquote(c.QNAME()), Hostname: host, Https: true}
		check.CheckHttpDigest = &CheckHttpDigest{CheckHttpBase: base, Digest: unquote(c.Cdigest().QNAME())}
		break
	}

	return &check, nil
}
