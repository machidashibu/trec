package presenter

import (
	"trec/internal/domain"
)

type LookupCollapsedReporter struct {
	table tableReporter
	timef *DurationFormatter
}

func NewLookupCollapsedReporter(table tableReporter, timef *DurationFormatter) *LookupCollapsedReporter {
	return &LookupCollapsedReporter{
		table: table,
		timef: timef,
	}
}

func (r LookupCollapsedReporter) Report(list domain.CollapsedTestList) {
	r.table.Header("name", "duration", "count")
	for _, item := range list {
		r.table.Row(item.Name(), r.timef.String(item.Duration()), item.Count())
	}
	r.table.Save()
}
