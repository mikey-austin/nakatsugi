package config

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
	Items []string
}


