package domain

import (
	"errors"
	"time"
)

var (
	ErrorInvalidConfig          = errors.New("invalid config")
	ErrorInvalidTestNamePattern = errors.New("invalid testname pattern")
	ErrorInvalidOperation       = errors.New("invalid operation")
	ErrorUnknownFormat          = errors.New("unknown format")
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

type CollaposedTest interface {
	Name() string
	Duration() time.Duration
	Count() int
}

type CollapsedTestList []CollaposedTest

type LookupStyle string

const (
	LookupFull     LookupStyle = "full"
	LookupSimple   LookupStyle = "simple"
	LookupCollapse LookupStyle = "collapse"
)

type TableFormat string

const (
	TableCsv        TableFormat = "csv"
	TableJson       TableFormat = "json"
	TableMarkdown   TableFormat = "markdown"
	TablePrettyJson TableFormat = "pretty json"
	TableText       TableFormat = "text"
)

type RecordId uint64

type OrderBy string

const (
	OrderByID       OrderBy = "id"
	OrderByName     OrderBy = "name"
	OrderByStart    OrderBy = "start_time"
	OrderByEnd      OrderBy = "end_time"
	OrderByResult   OrderBy = "result"
	OrderByDuration OrderBy = "duration"
	OrderByAsc      OrderBy = "ASC"
	OrderByDesc     OrderBy = "DESC"
)

type Filter interface {
	Today() bool
	LatestOnly() bool
}
