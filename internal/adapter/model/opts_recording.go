package model

import "time"

type RecordingOptions struct {
	timeFormat string
}

func NewRecordingOptions(timeFormat string) *RecordingOptions {
	return &RecordingOptions{
		timeFormat: timeFormat,
	}
}

// TODO: Fixed
func (ro RecordingOptions) Interval() time.Duration {
	return 1 * time.Second
}

func (ro RecordingOptions) TimeFormat() string {
	return ro.timeFormat
}
