package presenter

import (
	"fmt"
	"trec/internal/domain"
)

type lookupFormatter interface {
	String(id domain.RecordId, record domain.Test) string
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

func (r LookupReporter) Report(list domain.TestList) {
	for index := 0; index < list.Count(); index++ {
		test, id := list.Get(index)
		r.printer.Print(r.formatter.String(id, test))
	}
	r.printer.Print(fmt.Sprintf("%d items", list.Count()))
}
