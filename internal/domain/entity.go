package domain

import (
	"errors"
	"time"
)

var (
	ErrorInvalidConfig          = errors.New("invalid config")
	ErrorInvalidTestNamePattern = errors.New("invalid testname pattern")
)

type Test interface {
	Name() string
	StartTime() time.Time
	EndTime() time.Time
	Result() string
	Duration() time.Duration
}

type TestList interface {
	Count() int
	Get(index int) (Test, RecordId)
}

type RecordId uint64

type OrderBy string

const (
	OrderByID       OrderBy = "id"
	OrderByName     OrderBy = "name"
	OrderByStart    OrderBy = "start_time"
	OrderByEnd      OrderBy = "end_time"
	OrderByNote     OrderBy = "note"
	OrderByDuration OrderBy = "duration"
	OrderByAsc      OrderBy = "ASC"
	OrderByDesc     OrderBy = "DESC"
)

type Filter interface {
	Today() bool
	LatestOnly() bool
}
