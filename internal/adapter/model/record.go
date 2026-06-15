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

func (r Record) Label() string {
	return r.label
}

func (r Record) StartTime() time.Time {
	return r.startTime
}

func (r Record) EndTime() time.Time {
	return r.endTime
}

func (r Record) Note() string {
	return r.note
}
