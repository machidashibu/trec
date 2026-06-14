package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type databaseConfig interface {
	DBPath() string
}

type SQLite struct {
	db *gorm.DB
}

func OpenSQLite(config databaseConfig) (*SQLite, error) {
	db, err := gorm.Open(sqlite.Open(config.DBPath()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &SQLite{db: db}, nil
}

func (s *SQLite) DB() *gorm.DB {
	return s.db
}
