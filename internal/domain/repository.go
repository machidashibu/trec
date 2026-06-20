package domain

import (
	"time"
)

type TestResultRepository interface {
	Add(name string, start time.Time, end time.Time, memo string) (Test, error)
	GetAll(filter Filter) (TestList, error)
	Delete(id RecordId) error
}
