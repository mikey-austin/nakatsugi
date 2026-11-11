package config

import (
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
    example.com \
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
	if len(mytable.Items) != 3 {
		t.Errorf("expected %d mytable items, got %d", 3, len(mytable.Items))
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
}
