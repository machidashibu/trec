package core

import (
	"log/slog"
	"os"
	"regexp"
	"slices"
	"time"
	"trec/internal/core/logger"
	"trec/internal/domain"

	yaml "gopkg.in/yaml.v3"
)

type Mode string

const (
	ModeRecording Mode = "recording"
	ModeLookup    Mode = "lookup"
	ModeHelp      Mode = "help"
	ModeUnknown   Mode = "unknown"
)

// Config holds the application configuration settings.
// It is a stub code.
type Config struct {
	Mode      Mode            `yaml:"mode"`
	Recording RecordingConfig `yaml:"recording"`
	Lookup    LookupConfig    `yaml:"lookup"`
}

type RecordingConfig struct {
	DefaultLabel      string `yaml:"label"`
	ValidationPattern string `yaml:"validation"`
	DefaultTimeformat string `yaml:"time_format"`
}

type LookupConfig struct {
	DefaultOrder      domain.OrderBy     `yaml:"order"`
	DefaultFormat     string             `yaml:"format"`
	DefaultTimeformat string             `yaml:"time_format"`
	DefaultFilter     LookupFilterConfig `yaml:"filter"`
}

type LookupFilterConfig struct {
	StartTimeToday     bool `yaml:"today"`
	LatestOnlyPerLabel bool `yaml:"latest_only"`
}

func (c *Config) Read(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(f, &c)
}

func (c *Config) ParseArgs(args []string) error {
	// validate
	if len(args) < 1 {
		return domain.ErrorInvalidConfig
	}

	// parse mode
	switch args[0] {
	case "-l", "--lookup":
		c.Mode = ModeLookup
		return c.ParseLookupOptions(args[1:])
	case "-h", "--help":
		c.Mode = ModeHelp
	default: // default: recoding mode
		c.Mode = ModeRecording
		if args[0] == "-r" || args[0] == "--recording" {
			// specified mode
			return c.ParseRecordingOptions(args[1:])
		} else {
			// not specified mode
			return c.ParseRecordingOptions(args)
		}
	}

	return nil
}

func (c *Config) ParseLookupOptions(args []string) error {
	order_col := domain.OrderByID // TODO
	order_dir := domain.OrderByAsc
	if slices.Contains(args, "--desc") {
		order_dir = domain.OrderByDesc
	}
	c.Lookup.DefaultOrder = order_col + " " + order_dir
	if slices.Contains(args, "--format-full") {
		c.Lookup.DefaultFormat = "full"
	}
	if slices.Contains(args, "--no-filter") {
		// clear all filter
		c.Lookup.DefaultFilter.StartTimeToday = false
		c.Lookup.DefaultFilter.LatestOnlyPerLabel = false
	} else {
		// parse filter options
		if slices.Contains(args, "--all-days") {
			c.Lookup.DefaultFilter.StartTimeToday = false
		} else if slices.Contains(args, "--today") {
			c.Lookup.DefaultFilter.StartTimeToday = true
		}
		if slices.Contains(args, "--latest-only") {
			c.Lookup.DefaultFilter.LatestOnlyPerLabel = true
		}
	}
	return nil
}

func (c *Config) ParseRecordingOptions(args []string) error {
	if len(args) < 1 {
		return domain.ErrorInvalidConfig
	}
	c.Recording.DefaultLabel = args[0]
	if c.Recording.ValidationPattern != "" {
		matched, err := regexp.MatchString(c.Recording.ValidationPattern, c.Recording.DefaultLabel)
		if err != nil {
			return logger.Error("Config", "label validattion pattern error", err, "pattern", c.Recording.ValidationPattern, "label", c.Recording.DefaultLabel)
		}
		if !matched {
			return logger.Error("config", "label validation error", domain.ErrorInvalidLabelPattern, "pattern", c.Recording.ValidationPattern, "label", c.Recording.DefaultLabel)
		}
	}
	return nil
}

func (c Config) LogPath() string {
	return "trec.log"
}

func (c Config) LogLevel() slog.Level {
	return slog.LevelDebug
}

func (c Config) LogIsOverwrite() bool {
	return true
}

func (c Config) DBPath() string {
	return "trec.db"
}

func (c Config) Interval() time.Duration {
	return 1 * time.Second
}

func (c Config) AppMode() Mode {
	return c.Mode
}

func (c Config) Label() string {
	return c.Recording.DefaultLabel
}

func (c Config) RecordingTimeformat() string {
	return c.Recording.DefaultTimeformat
}

func (c Config) LookupOrder() domain.OrderBy {
	return domain.OrderBy(c.Lookup.DefaultOrder)
}

func (c Config) LookupFormat() string {
	if c.Lookup.DefaultFormat == "" {
		return "simple"
	} else {
		return c.Lookup.DefaultFormat
	}
}

func (c Config) LookupTimeFormat() string {
	return c.Lookup.DefaultTimeformat
}

func (c Config) LookupFilter() domain.Filter {
	return c.Lookup.DefaultFilter
}

func (c LookupFilterConfig) Today() bool {
	return c.StartTimeToday
}

func (c LookupFilterConfig) LatestOnly() bool {
	return c.LatestOnlyPerLabel
}
