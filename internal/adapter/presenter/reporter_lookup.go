package presenter

import (
	"fmt"
	"maps"
	"slices"
	"trec/internal/domain"
)

type lookupFormatter interface {
	String(id domain.OrderID, record domain.Record) string
}

type LookupReporter struct {
	printer   reportPrinter
	formatter lookupFormatter
}

func NewLookupReporter(printer reportPrinter, formatter lookupFormatter) *LookupReporter {
	return &LookupReporter{
		printer:   printer,
		formatter: formatter,
	}
}

func (r LookupReporter) Report(list domain.RecordList) {
	for _, id := range slices.Sorted(maps.Keys(list)) {
		r.printer.Print(r.formatter.String(domain.OrderID(id), list[id]))
	}
	r.printer.Print(fmt.Sprintf("%d items", len(list)))
}
