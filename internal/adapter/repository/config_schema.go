package repository

import (
	"log/slog"
	"trec/internal/domain"
)

type Config struct {
	Log       LogConfig       `yaml:"yaml:"log""`
	Recording RecordingConfig `yaml:"recording"`
	Lookup    LookupConfig    `yaml:"lookup"`
}

type LogConfig struct {
	Path      string     `yaml:"path"`
	Level     slog.Level `yaml:"level"`
	Overwrite bool       `yaml:"overwrite"`
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
