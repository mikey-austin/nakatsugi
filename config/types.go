package config

import (
	"net"
	"time"
)

type Config struct {
	Interval  int
	Tables    map[string]*Table
	Redirects map[string]*Redirect

	// TODO:
	// - Relays
	// - Routers
	// - Protocols
}

// Table types
type Table struct {
	Name     string
	Disabled bool
	Items    []*TItem
}

type TItemType int

const (
	IPv4Host TItemType = iota
	IPv6Host
	Hostname
)

type TItem struct {
	Type     TItemType
	Hostname string
	IP       net.IP

	TTL    time.Duration
	Parent uint
	Prio   uint
	Retry  uint
}

// Common types
type Proto string

const (
	ProtoTcp Proto = "tcp"
	ProtoUdp Proto = "udp"
)

type Listen struct {
	Address   net.IP
	Proto     Proto
	MinPort   uint16
	MaxPort   uint16
	Interface string
}

type Destination struct {
	ToTable  *Table
	Port     uint16
	Check    *Check
	Interval int
	Timeout  time.Duration

	// TODO: add mode here
}

type Forward struct {
	*Destination
}

type Route struct {
	*Destination
}

// Redirect types
type Redirect struct {
	Name           string
	Disabled       bool
	SessionTimeout time.Duration
	StickyAddress  bool
	Listens        []*Listen
	Forwards       []*Forward
	Routes         []*Route
}

// Check types
type CheckType int

const (
	CheckTypeHttpCode CheckType = iota
	CheckTypeHttpDigest
)

type Check struct {
	Type            CheckType
	CheckHttpCode   *CheckHttpCode
	CheckHttpDigest *CheckHttpDigest
}

type CheckHttpBase struct {
	Path     string
	Hostname string
	Https    bool
}

type CheckHttpCode struct {
	CheckHttpBase
	Code uint
}

type CheckHttpDigest struct {
	CheckHttpBase
	Digest string
}
