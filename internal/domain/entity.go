package domain

import (
	"errors"
	"time"
)

var (
	ErrorInvalidConfig = errors.New("invalid config")
)

type Record interface {
	Label() string
	StartTime() time.Time
	EndTime() time.Time
	Note() string
}

type RecordList interface {
	Count() int
	Get(index int) (uint64, Record)
}

type OrderBy string

const (
	OrderByID       OrderBy = "id"
	OrderByLabel    OrderBy = "label"
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
