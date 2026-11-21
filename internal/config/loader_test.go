package config

import (
	"net"
	"reflect"
	"strings"
	"testing"
	"time"
)

// TODO: add more test cases for error handling paths

func TestSimpleConfigLoad(t *testing.T) {
	c := `
myvar = test.com

interval 39 # seconds
timeout 210 # millis
socket "/tmp/my.sock"

# A normal table over multiple lines
table <mytable> { \
    x.org, \
    $myvar, \
}

# A disabled table
table <dtable> disable { $myvar }
`
	config, error := NewConfigFromReader(strings.NewReader(c))
	if error != nil {
		t.Errorf("could not load config: %v", error)
	}

	expectedSocket := "/tmp/my.sock"
	if config.SocketPath != expectedSocket {
		t.Errorf("expected %v, got %v", expectedSocket, config.SocketPath)
	}

	expectedInterval := time.Duration(39) * time.Second
	if config.Interval != expectedInterval {
		t.Errorf("expected %v, got %v", expectedInterval, config.Interval)
	}

	expectedTimeout := time.Duration(210) * time.Millisecond
	if config.Timeout != expectedTimeout {
		t.Errorf("expected %v, got %v", expectedTimeout, config.Timeout)
	}

	mytable := config.Tables["mytable"]
	if mytable == nil {
		t.Error("mytable is nil")
	}
	if mytable.Name != "mytable" {
		t.Errorf("table name %v expected, got %v", "mytable", mytable.Name)
	}
	if mytable.Disabled {
		t.Errorf("expected mytable to not be disabled")
	}
	if len(mytable.Items) != 2 {
		t.Errorf("expected %d mytable items, got %d", 2, len(mytable.Items))
	}
	if mytable.Items[0].Hostname != "x.org" || mytable.Items[1].Hostname != "test.com" {
		t.Errorf("mytable item unexpected: %v", mytable.Items)
	}

	dtable := config.Tables["dtable"]
	if dtable == nil {
		t.Error("dtable is nil")
	}
	if !dtable.Disabled {
		t.Errorf("expected dtable to be disabled")
	}
	if len(dtable.Items) != 1 {
		t.Errorf("expected %d dtable items, got %d", 1, len(dtable.Items))
	}
	if dtable.Items[0].Hostname != "test.com" {
		t.Errorf("dtable item unexpected: %v", dtable.Items)
	}
}

func TestTableItemOpts(t *testing.T) {
	c := `
table <mytable> { \
    x.org ip ttl 12 parent 5 priority 23 retry 3, \
    10.23.4.1 parent 7 ip ttl 21 retry 5 priority 27, \
    ::ffff:c000:0280 ip ttl 34 parent 8 retry 8 priority 34 }
`
	config, error := NewConfigFromReader(strings.NewReader(c))
	if error != nil {
		t.Errorf("could not load config: %v", error)
	}

	mytable, ok := config.Tables["mytable"]
	if !ok {
		t.Error("mytable is nil")
		return
	}
	if len(mytable.Items) != 3 {
		t.Errorf("expected %d mytable items, got %d", 2, len(mytable.Items))
	}

	expected := []TableItem{
		TableItem{Type: Hostname, Hostname: "x.org", TTL: time.Duration(12) * time.Second, Parent: 5, Prio: 23, Retry: 3},
		TableItem{Type: IPv4Host, IP: net.ParseIP("10.23.4.1"), TTL: time.Duration(21) * time.Second, Parent: 7, Prio: 27, Retry: 5},
		TableItem{Type: IPv6Host, IP: net.ParseIP("::ffff:c000:0280"), TTL: time.Duration(34) * time.Second, Parent: 8, Prio: 34, Retry: 8},
	}
	for i, e := range expected {
		if !reflect.DeepEqual(*(mytable.Items[i]), e) {
			t.Errorf("mytable.Items[%d]: expected %v, got %v", i, e, *mytable.Items[i])
		}
	}
}

func TestRedirect(t *testing.T) {
	c := `
redirect "udprange" {
    disable
    sticky-address
    session timeout 13
    listen on 10.23.1.3 proto udp port 333:335 interface em0
}

redirect "tcpsingle" {
    listen on 10.23.1.2 port 555
}
`
	config, error := NewConfigFromReader(strings.NewReader(c))
	if error != nil {
		t.Errorf("could not load config: %v", error)
	}

	udprange, ok := config.Redirects["udprange"]
	if !ok {
		t.Error("udprange is nil")
		return
	}
	tcpsingle, ok := config.Redirects["tcpsingle"]
	if !ok {
		t.Error("tcpsingle is nil")
		return
	}

	got := []*Redirect{udprange, tcpsingle}
	expected := []*Redirect{
		&Redirect{
			Name:           "udprange",
			Disabled:       true,
			StickyAddress:  true,
			SessionTimeout: time.Duration(13) * time.Second,
			Listens: []*Listen{
				&Listen{Address: net.ParseIP("10.23.1.3"), Proto: ProtoUdp, MinPort: 333, MaxPort: 335, Interface: "em0"},
			}},
		&Redirect{
			Name:           "tcpsingle",
			Disabled:       false,                            // Default
			StickyAddress:  false,                            // Default
			SessionTimeout: time.Duration(600) * time.Second, // Default
			Listens: []*Listen{
				&Listen{Address: net.ParseIP("10.23.1.2"), Proto: ProtoTcp, MinPort: 555, MaxPort: 555, Interface: ""},
			}},
	}

	for i, e := range expected {
		g := got[i]
		if !reflect.DeepEqual(&e, &g) {
			t.Errorf("TestRedirect(%d): expected %v, got %v", i, e, g)
		}
	}
}

func TestRedirectForwardsAndRoutes(t *testing.T) {
	c := `
interval 5
table <t1> { 10.23.1.1 }
redirect "r1" {
    listen on 10.23.1.2 port 555

    # Test forward rule extraction
    forward to <t1> port 777 \
        check http "/" host example.com code 200 \
        mode hash "myhashvalue" \
        interval 10 \
        timeout 25

    # Should default to the last listen port
    forward to <t1> \
        timeout 27 \
        interval 15 \
        mode least-states \
        check http "/health" digest "a9993e36476816aba3e25717850c26c9cd0d89d"

    forward to <t1> \
        timeout 27 \
        interval 15 \
        mode loadbalance "mybalancer" \
        check https "/" digest "a9993e36476816aba3e25717850c26c9cd0d89d"

    forward to <t1> \
        mode random \
        check https "/" host test.com code 202

    forward to <t1> \
        mode roundrobin \
        check https "/" host test.com code 204

    forward to <t1> \
        mode source-hash "mysourcehash" \
        check https "/" host test.com code 300

    # Now the route rule versions of the above rules
    route to <t1> port 777 \
        check http "/" host example.com code 200 \
        mode hash "myhashvalue" \
        interval 10 \
        timeout 25

    # Should default to the last listen port
    route to <t1> \
        timeout 27 \
        interval 15 \
        mode least-states \
        check http "/health" digest "a9993e36476816aba3e25717850c26c9cd0d89d"

    route to <t1> \
        timeout 27 \
        interval 15 \
        mode loadbalance "mybalancer" \
        check https "/" digest "a9993e36476816aba3e25717850c26c9cd0d89d"

    route to <t1> \
        mode random \
        check https "/" host test.com code 202

    route to <t1> \
        mode roundrobin \
        check https "/" host test.com code 204

    route to <t1> \
        mode source-hash "mysourcehash" \
        check https "/" host test.com code 300
}
`
	config, error := NewConfigFromReader(strings.NewReader(c))
	if error != nil {
		t.Errorf("could not load config: %v", error)
	}

	r1, ok := config.Redirects["r1"]
	if !ok {
		t.Error("udprange is nil")
		return
	}

	table := Table{
		Name: "t1",
		Items: []*TableItem{
			&TableItem{Type: IPv4Host, IP: net.ParseIP("10.23.1.1")},
		},
	}
	expected := []*Destination{
		&Destination{
			ToTable: &table,
			Port:    777,
			Mode:    &ModeHash{hash: "myhashvalue"},
			Check: &Check{
				Type: CheckTypeHttpCode,

				CheckHttpCode: &CheckHttpCode{
					CheckHttpBase: CheckHttpBase{Path: "/", Hostname: "example.com", Https: false},
					Code:          200,
				}},
		},
		&Destination{
			ToTable: &table,
			Port:    555, // Defaults to the listen port defined above it
			Mode:    &ModeLeastStates{},
			Check: &Check{
				Type: CheckTypeHttpDigest,
				CheckHttpDigest: &CheckHttpDigest{
					CheckHttpBase: CheckHttpBase{Path: "/health", Hostname: "", Https: false},
					Digest:        "a9993e36476816aba3e25717850c26c9cd0d89d",
				}},
		},
		&Destination{
			ToTable: &table,
			Port:    555, // Defaults to the listen port defined above it
			Mode:    &ModeLoadBalance{name: "mybalancer"},
			Check: &Check{
				Type: CheckTypeHttpDigest,
				CheckHttpDigest: &CheckHttpDigest{
					CheckHttpBase: CheckHttpBase{Path: "/", Hostname: "", Https: true},
					Digest:        "a9993e36476816aba3e25717850c26c9cd0d89d",
				}},
		},
		&Destination{
			ToTable: &table,
			Port:    555, // Defaults to the listen port defined above it
			Mode:    &ModeRandom{},
			Check: &Check{
				Type: CheckTypeHttpCode,
				CheckHttpCode: &CheckHttpCode{
					CheckHttpBase: CheckHttpBase{Path: "/", Hostname: "test.com", Https: true},
					Code:          202,
				}},
		},
		&Destination{
			ToTable: &table,
			Port:    555, // Defaults to the listen port defined above it
			Mode:    &ModeRoundRobin{},
			Check: &Check{
				Type: CheckTypeHttpCode,
				CheckHttpCode: &CheckHttpCode{
					CheckHttpBase: CheckHttpBase{Path: "/", Hostname: "test.com", Https: true},
					Code:          204,
				}},
		},
		&Destination{
			ToTable: &table,
			Port:    555, // Defaults to the listen port defined above it
			Mode:    &ModeSourceHash{hash: "mysourcehash"},
			Check: &Check{
				Type: CheckTypeHttpCode,
				CheckHttpCode: &CheckHttpCode{
					CheckHttpBase: CheckHttpBase{Path: "/", Hostname: "test.com", Https: true},
					Code:          300,
				}},
		},
	}

	for i, e := range expected {
		g := r1.Forwards[i]
		if e.Port != g.Port {
			t.Errorf("TestForward(%d).Port: expected %v, got %v", i, e.Port, g.Port)
		}
		if !reflect.DeepEqual(e.ToTable, g.ToTable) {
			t.Errorf("TestForward(%d).ToTable: expected %v, got %v", i, e.ToTable, g.ToTable)
		}
		if !reflect.DeepEqual(e.Check, g.Check) {
			t.Errorf("TestForward(%d).Check: expected %v, got %v", i, e.Check, g.Check)
		}
		if !reflect.DeepEqual(e.Mode, g.Mode) {
			t.Errorf("TestForward(%d).Mode: expected %v, got %v", i, e.Mode, g.Mode)
		}
	}

	for i, e := range expected {
		g := r1.Routes[i]
		if e.Port != g.Port {
			t.Errorf("TestRoute(%d).Port: expected %v, got %v", i, e.Port, g.Port)
		}
		if !reflect.DeepEqual(e.ToTable, g.ToTable) {
			t.Errorf("TestRoute(%d).ToTable: expected %v, got %v", i, e.ToTable, g.ToTable)
		}
		if !reflect.DeepEqual(e.Check, g.Check) {
			t.Errorf("TestRoute(%d).Check: expected %v, got %v", i, e.Check, g.Check)
		}
		if !reflect.DeepEqual(e.Mode, g.Mode) {
			t.Errorf("TestRoute(%d).Mode: expected %v, got %v", i, e.Mode, g.Mode)
		}
	}
}
