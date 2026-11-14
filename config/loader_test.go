package config

import (
	"net"
	"reflect"
	"strings"
	"testing"
)

func TestSimpleConfigLoad(t *testing.T) {
	c := `
myvar = test.com

interval 39

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

	if config.Interval != 39 {
		t.Errorf("expected %v, got %v", 39, config.Interval)
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

	expected := []TItem{
		TItem{Type: Hostname, Hostname: "x.org", TTL: 12, Parent: 5, Prio: 23, Retry: 3},
		TItem{Type: IPv4Host, IP: net.ParseIP("10.23.4.1"), TTL: 21, Parent: 7, Prio: 27, Retry: 5},
		TItem{Type: IPv6Host, IP: net.ParseIP("::ffff:c000:0280"), TTL: 34, Parent: 8, Prio: 34, Retry: 8},
	}
	for i, e := range expected {
		if !reflect.DeepEqual(*(mytable.Items[i]), e) {
			t.Errorf("mytable.Items[%d]: expected %v, got %v", i, e, *mytable.Items[i])
		}
	}
}
