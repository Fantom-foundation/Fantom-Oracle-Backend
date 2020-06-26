// Package main implements the entry point into the Watchdog app.
package main

import "oracle-watchdog/internal/config"

// main initializes the Watchdog server and starts it when ready.
func main() {
	// read the configuration
	cfg := config.Config()
}
