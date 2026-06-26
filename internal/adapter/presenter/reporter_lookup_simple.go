package presenter

import (
	"fmt"
	"trec/internal/domain"
)

type lookupSimpleOutput interface {
	Print(text string)
}

type LookupSimpleReporter struct {
	out lookupSimpleOutput
	tf  *DurationFormatter
}

func NewLookupSimpleReporter(out lookupSimpleOutput, tf *DurationFormatter) *LookupSimpleReporter {
	return &LookupSimpleReporter{
		out: out,
		tf:  tf,
	}
}

func (r LookupSimpleReporter) Report(list domain.TestList) {
	for index := 0; index < list.Count(); index++ {
		test, _ := list.Get(index)
		r.out.Print(fmt.Sprintf("%s %s %s", test.Name(), test.Result(), r.tf.String(test.Duration())))
	}
	r.out.Print(fmt.Sprintf("%d items", list.Count()))
}
