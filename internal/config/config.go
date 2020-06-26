// Package main implements the entry point into the Watchdog app.
package config

import "flag"

// Cfg declares the app configuration options container.
type Cfg struct {
	// NodeRpc represents the path to Lachesis RPC/IPC communication interface.
	NodeRpc string

	// ContractsRepoPath is the path to watched contracts configuration.
	ContractsRepoPath string
}

// Config loads and parses the configuration from the app flags
// and returns the configuration structure.
func Config() Cfg {
	// make the container
	var cfg Cfg

	// define flags
	flag.StringVar(&cfg.NodeRpc, "rpc", "/var/opera/lachesis/data/lachesis.ipc", "path to the Lachesis IPC/RPC interface")
	flag.StringVar(&cfg.ContractsRepoPath, "repo", "/var/opera/watchdog/contracts.json", "path to the contracts configuration file")
	flag.Parse()

	return cfg
}
