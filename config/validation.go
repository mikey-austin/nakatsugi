package config

import (
	"math"
	"regexp"
	"slices"
)

var nameRegexp = regexp.MustCompile("^[a-z0-9_]+$")
var hostnameRegexp = regexp.MustCompile("^([a-z0-9-]+)(\\.[a-z0-9-]+)+$")
var reservedWords = []string{
	"table", "relay", "timeout",
}

func IsValidPort(port uint64) bool {
	return port >= 0 && port <= math.MaxUint16
}

func IsValidName(name string) bool {
	return name != "" && nameRegexp.Match([]byte(name))
}

func IsValidVarName(name string) bool {
	return name != "" && nameRegexp.Match([]byte(name)) && !slices.Contains(reservedWords, name)
}

func IsValidHostname(name string) bool {
	return name != "" && hostnameRegexp.Match([]byte(name))
}
