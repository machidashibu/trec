package presenter

import (
	"time"
	"trec/internal/domain"
)

type LookupFullReporter struct {
	table tableReporter
	timef *DurationFormatter
}

func NewLookupFullReporter(table tableReporter, timef *DurationFormatter) *LookupFullReporter {
	return &LookupFullReporter{
		table: table,
		timef: timef,
	}
}

func (r LookupFullReporter) Report(list domain.TestList) {
	r.table.Header("id", "name", "result", "start_time", "end_time", "duration")
	for index := 0; index < list.Count(); index++ {
		test, id := list.Get(index)
		r.table.Row(
			id,
			test.Name(),
			test.Result(),
			test.StartTime().Format(time.DateTime),
			test.EndTime().Format(time.DateTime),
			r.timef.String(test.Duration()),
		)
	}
	r.table.Save()
}
