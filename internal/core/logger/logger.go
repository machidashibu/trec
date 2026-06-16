package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type loggerConfig interface {
	LogPath() string
	LogLevel() slog.Level
	LogIsOverwrite() bool
}

func ApplyConfig(config loggerConfig) error {
	var logw io.Writer = io.Discard // default: no logging
	if config.LogPath() != "" {
		if config.LogIsOverwrite() {
			// Create or overwrite the log file
			f, err := os.Create(config.LogPath())
			if err != nil {
				return fmt.Errorf("failed to create log file: %w", err)
			}
			logw = f
		} else {
			// Open the log file for appending
			f, err := os.OpenFile(config.LogPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return fmt.Errorf("failed to open log file: %w", err)
			}
			logw = f
		}
	}

	// Create a new logger with the specified log level and output
	slog.SetDefault(slog.New(slog.NewTextHandler(logw, &slog.HandlerOptions{Level: config.LogLevel()})))

	return nil
}

func Error(tag string, message string, err error, args ...any) error {
	errArgs := make([]any, 0, len(args)+4)
	errArgs = append(errArgs, "error", err)
	errArgs = append(errArgs, "tag", tag)
	errArgs = append(errArgs, args...)
	slog.Error(tag+": "+message, errArgs...)
	return err
}
