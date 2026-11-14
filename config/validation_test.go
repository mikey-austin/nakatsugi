package config

import (
	"testing"
)

func TestTableNameValidation(t *testing.T) {
	cases := []struct {
		I string
		R bool
	}{
		{I: "mytable", R: true},
		{I: "Mytable", R: false},
		{I: "my-table", R: false},
		{I: "my_table", R: true},
		{I: "my_table1", R: true},
		{I: "my table1", R: false},
		{I: "", R: false},
	}
	for i, c := range cases {
		if got := IsValidTableName(c.I); got != c.R {
			t.Errorf("IsValidTableName(%v)[%d]: expected %v, got %v", c.I, i, c.R, got)
		}
	}
}

func TestVarNameValidation(t *testing.T) {
	cases := []struct {
		I string
		R bool
	}{
		{I: "myvar", R: true},
		{I: "Myvar", R: false},
		{I: "my-var", R: false},
		{I: "my_var", R: true},
		{I: "my_var1", R: true},
		{I: "my var1", R: false},
		{I: "table", R: false},
		{I: "relay", R: false},
		{I: "timeout", R: false},
		{I: "my var1", R: false},
		{I: "", R: false},
	}
	for i, c := range cases {
		if got := IsValidVarName(c.I); got != c.R {
			t.Errorf("IsValidVarName(%v)[%d]: expected %v, got %v", c.I, i, c.R, got)
		}
	}
}

func TestHostnameValidation(t *testing.T) {
	cases := []struct {
		I string
		R bool
	}{
		{I: "myhost", R: false},
		{I: "Myhost.com", R: false},
		{I: "my-host.com", R: true},
		{I: "my_host98.com", R: false},
		{I: "myhost1.com", R: true},
		{I: "my.host.org", R: true},
		{I: ":table", R: false},
		{I: "", R: false},
	}
	for i, c := range cases {
		if got := IsValidHostname(c.I); got != c.R {
			t.Errorf("IsValidHostname(%v)[%d]: expected %v, got %v", c.I, i, c.R, got)
		}
	}
}
