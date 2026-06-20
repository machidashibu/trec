package model

import (
	"time"
	"trec/internal/domain"
)

type Test struct {
	name      string
	startTime time.Time
	endTime   time.Time
	result    string
}

func NewRecord(name string, startTime time.Time, endTime time.Time, result string) domain.Test {
	return Test{
		name:      name,
		startTime: startTime,
		endTime:   endTime,
		result:    result,
	}
}

func (t Test) Name() string {
	return t.name
}

func (t Test) StartTime() time.Time {
	return t.startTime
}

func (t Test) EndTime() time.Time {
	return t.endTime
}

func (t Test) Result() string {
	return t.result
}

func (t Test) Duration() time.Duration {
	return t.endTime.Sub(t.startTime)
}
