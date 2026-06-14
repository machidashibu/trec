package model

import (
	"time"
	"trec/internal/domain"
)

type Record struct {
	label     string
	startTime time.Time
	endTime   time.Time
	note      string
}

func NewRecord(label string, startTime time.Time, endTime time.Time, note string) domain.Record {
	return Record{
		label:     label,
		startTime: startTime,
		endTime:   endTime,
		note:      note,
	}
}

func (s Record) Label() string {
	return s.label
}

func (s Record) StartTime() time.Time {
	return s.startTime
}

func (s Record) EndTime() time.Time {
	return s.endTime
}

func (s Record) Note() string {
	return s.note
}
