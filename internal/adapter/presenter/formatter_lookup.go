package presenter

import (
	"fmt"
	"time"
	"trec/internal/domain"
)

type LookupFormatter struct {
	format     string
	timeFormat *DurationFormatter
}

func NewLookupFormatter(format, timeFormat string) *LookupFormatter {
	return &LookupFormatter{
		format:     format,
		timeFormat: NewDurationFormatter(timeFormat),
	}
}

func (f LookupFormatter) String(id domain.RecordId, test domain.Test) string {
	switch f.format {
	case "full":
		return fmt.Sprintf("[%d] %s %s %s %s %s",
			id,
			test.Name(),
			test.Result(),
			test.StartTime().Format(time.DateTime),
			test.EndTime().Format(time.DateTime),
			f.timeFormat.String(test.Duration()),
		)
	// case "simple":
	default:
		return fmt.Sprintf("%s %s %s", test.Name(), test.Result(), f.timeFormat.String(test.Duration()))
	}
}
