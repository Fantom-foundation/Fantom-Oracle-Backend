// Package logger defines leveled operation and state change recording service.
package logger

import (
	"github.com/op/go-logging"
	"oracle-watchdog/internal/config"
	"os"
)

// logAppName is the name of the application used in joined logs facility.
const logAppName = `OperaWatchdog`

// logFormat represents the predefined format of the log
const logFormat = `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s}%{color:reset} %{message}`

// ApiLogger defines extended logger with generic no-level logging option
type ApiLogger struct {
	logging.Logger
}

// Printf implements default non-leveled output.
// We assume the information is low in importance if passed to this function so we relay it to Debug level.
func (a ApiLogger) Printf(format string, args ...interface{}) {
	a.Debugf(format, args...)
}

// New provides pre-configured Logger with stderr output and leveled filtering.
// Modules are not supported at the moment, but may be added in the future
// to make the logging setup more granular.
func New(cfg *config.Cfg) Logger {
	// Prep the backend for exporting the log records
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	// Parse log format from configuration and apply it to the backend
	format := logging.MustStringFormatter(logFormat)
	fmtBackend := logging.NewBackendFormatter(backend, format)

	// Parse and apply the configured level on which the recording will be emitted
	level, err := logging.LogLevel(cfg.LogLevel)
	if err != nil {
		level = logging.INFO
	}
	lvlBackend := logging.AddModuleLevel(fmtBackend)
	lvlBackend.SetLevel(level, "")

	// assign the backend and return the new logger
	logging.SetBackend(lvlBackend)
	l := logging.MustGetLogger(logAppName)

	return &ApiLogger{*l}
}
