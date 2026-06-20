package core

import (
	"log/slog"
	"os"
	"trec/internal/domain"

	yaml "gopkg.in/yaml.v3"
)

type Mode string

const (
	ModeRecording Mode = "recording"
	ModeLookup    Mode = "lookup"
	ModeDelete    Mode = "delete"
	ModeHelp      Mode = "help"
	ModeUnknown   Mode = "unknown"
)

// Config holds the application configuration settings.
// It is a stub code.
type Config struct {
	Recording RecordingConfig `yaml:"recording"`
	Lookup    LookupConfig    `yaml:"lookup"`
}

type RecordingConfig struct {
	DefaultTestname   string `yaml:"testname"`
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
	StartTimeToday        bool `yaml:"today"`
	LatestOnlyPerTestname bool `yaml:"latest_only"`
}

func (c *Config) Read(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(f, &c)
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
