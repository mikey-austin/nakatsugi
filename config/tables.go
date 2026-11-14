package config

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/mikey-austin/nakatsugi/parser"
)

func StripTableName(raw string) string {
	name := strings.TrimLeft(raw, "<")
	name = strings.TrimRight(name, ">")
	return name

}

func NewTItemFromParseCtx(itemVal string, ctx parser.ITitemContext) (*TItem, error) {
	var item TItem
	ip := net.ParseIP(itemVal)
	if ip != nil {
		if strings.ContainsAny(itemVal, ":") {
			item.Type = IPv6Host
		} else {
			item.Type = IPv4Host
		}
		item.IP = ip
	} else {
		// We must then be a hostname
		if !IsValidHostname(itemVal) {
			return nil, fmt.Errorf("invalid hostname: %v", itemVal)
		}
		item.Type = Hostname
		item.Hostname = itemVal
	}

	// Now extract all table item options
	for _, topt := range ctx.AllTopts() {
		if ttl := topt.Toptipttl(); ttl != nil {
			parsedTtl, error := strconv.Atoi(ttl.NUMBER().GetText())
			if error != nil {
				return nil, fmt.Errorf("invalid ttl %q: %w", ttl.NUMBER().GetText(), error)
			}
			item.TTL = time.Duration(parsedTtl) * time.Second
		}

		if parent := topt.Toptparent(); parent != nil {
			parsedParent, error := strconv.ParseUint(parent.NUMBER().GetText(), 10, 32)
			if error != nil {
				return nil, fmt.Errorf("invalid parent: %v", parsedParent)
			}
			item.Parent = uint(parsedParent)
		}

		if priority := topt.Toptpriority(); priority != nil {
			parsedPriority, error := strconv.ParseUint(priority.NUMBER().GetText(), 10, 32)
			if error != nil {
				return nil, fmt.Errorf("invalid priority: %v", parsedPriority)
			}
			item.Prio = uint(parsedPriority)
		}

		if retry := topt.Toptretry(); retry != nil {
			parsedRetry, error := strconv.ParseUint(retry.NUMBER().GetText(), 10, 32)
			if error != nil {
				return nil, fmt.Errorf("invalid retry: %v", parsedRetry)
			}
			item.Retry = uint(parsedRetry)
		}
	}

	return &item, nil
}
