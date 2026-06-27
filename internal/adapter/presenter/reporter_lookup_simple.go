package presenter

import (
	"trec/internal/domain"
)

type LookupSimpleReporter struct {
	table tableReporter
	timef *DurationFormatter
}

func NewLookupSimpleReporter(table tableReporter, timef *DurationFormatter) *LookupSimpleReporter {
	return &LookupSimpleReporter{
		table: table,
		timef: timef,
	}
}

func (r LookupSimpleReporter) Report(list domain.TestList) {
	r.table.Header("name", "result", "duration")
	for index := 0; index < list.Count(); index++ {
		test, _ := list.Get(index)
		r.table.Row(test.Name(), test.Result(), r.timef.String(test.Duration()))
	}
	r.table.Save()
}
