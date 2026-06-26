package domain

import (
	"time"
)

// TestResultRepository provides methods to manage test result.
type TestResultRepository interface {
	Add(name string, start time.Time, end time.Time, result string) (Test, error)
	GetAll(filter Filter) (TestList, error)
	GetCollapsed(filter Filter) (CollapsedTestList, error)
	GetById(id RecordId) (Test, error)
	EditName(id RecordId, name string) error
	EditResult(id RecordId, result string) error
	Edit(id RecordId, name, result string) error
	Delete(id RecordId) error
}
