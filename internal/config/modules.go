// Package config implements the configuration of the Watchdog app.
package config

// ModuleConfig defines a general structure for a configuration
// of an oracle module supervised by the oracle supervisor.
type ModuleConfig struct {
	Type    string `json:"type"`
	CfgPath string `json:"cfg"`
}

// ModuleCatalog defines a structure of oracle modules configuration
// for the Supervisor.
type ModuleCatalog struct {
	Modules []ModuleConfig `json:"modules"`
}
