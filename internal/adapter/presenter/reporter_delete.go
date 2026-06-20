package presenter

import (
	"fmt"
	"trec/internal/domain"
)

type deleteOutput interface {
	Print(text string)
}

type DeleteReporter struct {
	out deleteOutput
}

func NewDeleteReporter(out deleteOutput) *DeleteReporter {
	return &DeleteReporter{
		out: out,
	}
}

func (r DeleteReporter) Report(id domain.RecordId) {
	r.out.Print(fmt.Sprintf("Deleted record. (%d)", id))
}
