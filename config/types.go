package config

import (
	"net"
)

type Config struct {
	Interval int
	Tables map[string]*Table

	// TODO:
	// - Redirects
	// - Relays
	// - Routers
	// - Protocols
}

type Table struct {
	Name string
	Disabled bool
	Items []*TItem
}

type TItemType int
const (
	IPv4Host TItemType = iota
	IPv6Host
	Hostname
)

type TItem struct {
	Type TItemType
	Hostname string
	IP net.IP

	TTL uint
	Parent uint
	Prio uint
	Retry uint
}


