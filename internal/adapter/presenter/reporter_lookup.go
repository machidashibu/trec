package presenter

import (
	"fmt"
	"trec/internal/domain"
)

type lookupOutput interface {
	Print(text string)
}

type lookupFormatter interface {
	String(id domain.RecordId, record domain.Test) string
}

type LookupReporter struct {
	out    lookupOutput
	format lookupFormatter
}

func NewLookupReporter(out lookupOutput, format lookupFormatter) *LookupReporter {
	return &LookupReporter{
		out:    out,
		format: format,
	}
}

func (r LookupReporter) Report(list domain.TestList) {
	for index := 0; index < list.Count(); index++ {
		test, id := list.Get(index)
		r.out.Print(r.format.String(id, test))
	}
	r.out.Print(fmt.Sprintf("%d items", list.Count()))
}
