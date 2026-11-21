package config

import (
	"fmt"
	"net"
	"time"
)

type Config struct {
	Interval   time.Duration
	Timeout    time.Duration
	Tables     map[string]*Table
	Redirects  map[string]*Redirect
	SocketPath string

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
	Interval time.Duration
	Timeout  time.Duration
	Check    *Check
	Mode     BalancerMode
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

func (c CheckType) String() string {
	switch c {
	case CheckTypeHttpCode:
		return "<check: http code>"
	case CheckTypeHttpDigest:
		return "<check: http digest>"
	default:
		return "<check: unknown>"
	}
}

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

func (c CheckHttpBase) String() string {
	return fmt.Sprintf("{path: '%s', hostname: '%s', https:%v}", c.Path, c.Hostname, c.Https)
}

type CheckHttpCode struct {
	CheckHttpBase
	Code uint
}

type CheckHttpDigest struct {
	CheckHttpBase
	Digest string
}

// Mode types
type BalancerMode interface{}
type ModeHash struct{ hash string }
type ModeLeastStates struct{}
type ModeLoadBalance struct{ name string }
type ModeRandom struct{}
type ModeRoundRobin struct{}
type ModeSourceHash struct{ hash string }
