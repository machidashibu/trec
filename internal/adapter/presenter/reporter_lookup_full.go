package presenter

import (
	"fmt"
	"time"
	"trec/internal/domain"
)

type lookupFullOutput interface {
	Print(text string)
}

type LookupFullReporter struct {
	out lookupFullOutput
	tf  *DurationFormatter
}

func NewLookupFullReporter(out lookupFullOutput, tf *DurationFormatter) *LookupFullReporter {
	return &LookupFullReporter{
		out: out,
		tf:  tf,
	}
}

func (r LookupFullReporter) Report(list domain.TestList) {
	for index := 0; index < list.Count(); index++ {
		test, id := list.Get(index)
		r.out.Print(fmt.Sprintf("[%d] %s %s %s %s %s",
			id,
			test.Name(),
			test.Result(),
			test.StartTime().Format(time.DateTime),
			test.EndTime().Format(time.DateTime),
			r.tf.String(test.Duration()),
		))
	}
	r.out.Print(fmt.Sprintf("%d items", list.Count()))
}
