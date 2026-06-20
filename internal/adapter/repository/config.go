package repository

import (
	"log/slog"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func (c *Config) Read(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(f, &c)
}

func (c Config) LogPath() string {
	if c.Log.Path == "" {
		return "trec.log"
	} else {
		return c.Log.Path
	}
}

func (c Config) LogLevel() slog.Level {
	return c.Log.Level
}

func (c Config) LogIsOverwrite() bool {
	return c.Log.Overwrite
}

func (c Config) DBPath() string {
	return "trec.db"
}
