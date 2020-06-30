// Package config implements the configuration of the Watchdog app.
package config

import "flag"

// Cfg declares the app configuration options container.
type Cfg struct {
	// NodeRpc represents the path to Lachesis RPC/IPC communication interface.
	NodeRpc string

	// ModulesConfigPath is the path to modules configuration JSON file.
	ModulesConfigPath string

	// LogLevel represents the level for log records of the oracle.
	LogLevel string
}

// Config loads and parses the configuration from the app flags
// and returns the configuration structure.
func New() Cfg {
	// make the container
	var cfg Cfg

	// define flags
	flag.StringVar(&cfg.NodeRpc, "rpc", "/var/opera/lachesis/data/lachesis.ipc", "path to the Lachesis IPC/RPC interface")
	flag.StringVar(&cfg.ModulesConfigPath, "cfg", "doc/config/modules.json", "path to the modules configuration JSON file")
	flag.StringVar(&cfg.LogLevel, "log", "NOTICE", "level of the log, use one of the [CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG]")
	flag.Parse()

	return cfg
}
