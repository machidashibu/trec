package core

import (
	"log/slog"
	"slices"
	"time"
	"trec/internal/domain"
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
	DefaultLabel string `yaml:"label"`
}

type LookupConfig struct {
	DefaultOrder  domain.OrderBy `yaml:"order"`
	DefaultFormat string         `yaml:"format"`
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
	c.Lookup.DefaultFormat = "simple" // default: simple
	if slices.Contains(args, "--format-full") {
		c.Lookup.DefaultFormat = "full"
	}
	return nil
}

func (c *Config) ParseRecordingOptions(args []string) error {
	if len(args) < 1 {
		return domain.ErrorInvalidConfig
	}
	c.Recording.DefaultLabel = args[0]
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

func (c Config) LookupOrder() domain.OrderBy {
	return domain.OrderBy(c.Lookup.DefaultOrder)
}

func (c Config) LookupFormat() string {
	return c.Lookup.DefaultFormat
}
