package presenter

import (
	"fmt"
	"trec/internal/domain"
)

type lookupFormatter interface {
	String(id uint64, record domain.Record) string
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
	for index := 0; index < list.Count(); index++ {
		id, record := list.Get(index)
		r.printer.Print(r.formatter.String(id, record))
	}
	r.printer.Print(fmt.Sprintf("%d items", list.Count()))
}
