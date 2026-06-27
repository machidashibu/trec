package repository

import (
	"log/slog"
	"trec/internal/domain"
)

type Config struct {
	Log       LogConfig       `yaml:"log"`
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
	DefaultStyle      domain.LookupStyle `yaml:"style"`
	DefaultFormat     domain.TableFormat `yaml:"format"`
	DefaultTimeformat string             `yaml:"time_format"`
	DefaultOrder      domain.OrderBy     `yaml:"order"`
	DefaultFilter     LookupFilterConfig `yaml:"filter"`
}

type LookupFilterConfig struct {
	StartTimeToday        bool `yaml:"today"`
	LatestOnlyPerTestname bool `yaml:"latest_only"`
	Collapse              bool `yaml:"collapse"`
}
