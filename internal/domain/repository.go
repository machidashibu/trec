package domain

import (
	"time"
)

type RecordRepository interface {
	Add(label string, start time.Time, end time.Time, memo string) (Record, error)
	GetAll(order OrderBy) (RecordList, error)
}
