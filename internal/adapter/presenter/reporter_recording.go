package presenter

import (
	"fmt"
	"time"
)

type recordingOutput interface {
	PrintLine(text string)
	Print(text string)
}

type recordingTimeFormatter interface {
	String(d time.Duration) string
}

type RecordingReporter struct {
	out    recordingOutput
	format recordingTimeFormatter
}

func NewRecordingReporter(out recordingOutput, format recordingTimeFormatter) *RecordingReporter {
	return &RecordingReporter{
		out:    out,
		format: format,
	}
}

func (r RecordingReporter) Start() {
	r.Ticked(0)
}

func (r RecordingReporter) Stop() {
	r.out.Print("")
}

func (r RecordingReporter) Ticked(d time.Duration) {
	r.out.PrintLine(fmt.Sprintf("Recording... %s", r.format.String(d)))
}

func (r RecordingReporter) Completed() {
	r.out.Print("Recorded.")
}
