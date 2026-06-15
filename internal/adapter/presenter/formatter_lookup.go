package presenter

import (
	"fmt"
	"time"
	"trec/internal/domain"
)

type LookupFormatter struct {
	format string
}

func NewLookupFormatter(format string) *LookupFormatter {
	return &LookupFormatter{
		format: format,
	}
}

func (f LookupFormatter) String(id uint64, record domain.Record) string {
	d := record.EndTime().Sub(record.StartTime())
	switch f.format {
	case "full":
		return fmt.Sprintf("[%d] %s %s %s %s %s",
			id,
			record.Label(),
			record.Note(),
			record.StartTime().Format(time.DateTime),
			record.EndTime().Format(time.DateTime),
			ToDurationString(d),
		)
	// case "simple":
	default:
		return fmt.Sprintf("%s %s %s", record.Label(), record.Note(), ToDurationString(d))
	}
}
