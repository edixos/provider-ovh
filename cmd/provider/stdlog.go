/*
Copyright 2026 The Crossplane Authors.
*/

package main

import (
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
)

// stdLogWriter adapts standard library log output, as emitted by the
// Terraform provider ("[DEBUG] message"), to a crossplane logging.Logger.
// [TRACE] and [DEBUG] lines are logged at debug level, everything else at
// info level with the original level kept as a key.
type stdLogWriter struct {
	log logging.Logger
}

func (w stdLogWriter) Write(p []byte) (int, error) {
	level, msg := parseStdLogLine(string(p))
	switch level {
	case "TRACE", "DEBUG":
		w.log.Debug(msg)
	case "":
		w.log.Info(msg)
	default:
		w.log.Info(msg, "level", strings.ToLower(level))
	}
	return len(p), nil
}

// parseStdLogLine splits a "[LEVEL] message" line into its level and message.
// Lines without a level prefix return an empty level.
func parseStdLogLine(line string) (string, string) {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "[") {
		if end := strings.Index(line, "]"); end > 1 {
			return strings.ToUpper(line[1:end]), strings.TrimSpace(line[end+1:])
		}
	}
	return "", line
}
