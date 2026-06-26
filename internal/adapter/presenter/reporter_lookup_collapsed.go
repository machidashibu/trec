package presenter

import (
	"fmt"
	"trec/internal/domain"
)

type lookupCollapsedOutput interface {
	Print(text string)
}

type LookupCollapsedReporter struct {
	out lookupCollapsedOutput
	tf  *DurationFormatter
}

func NewLookupCollapsedReporter(out lookupCollapsedOutput, tf *DurationFormatter) *LookupCollapsedReporter {
	return &LookupCollapsedReporter{
		out: out,
		tf:  tf,
	}
}

func (r LookupCollapsedReporter) Report(list domain.CollapsedTestList) {
	for _, item := range list {
		r.out.Print(fmt.Sprintf("%s %s (%d times)", item.Name(), r.tf.String(item.Duration()), item.Count()))
	}
	r.out.Print(fmt.Sprintf("%d items", len(list)))
}
