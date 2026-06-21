package presenter

import "fmt"

type editOutput interface {
	Print(text string)
}

type EditReporter struct {
	out editOutput
}

func NewEditReporter(out editOutput) *EditReporter {
	return &EditReporter{
		out: out,
	}
}

func (r EditReporter) ReportUpdatedName(name string) {
	r.out.Print(fmt.Sprintf("Updated name (%s)", name))
}

func (r EditReporter) ReportUpdatedResult(result string) {
	r.out.Print(fmt.Sprintf("Updated result (%s)", result))
}

func (r EditReporter) ReportNoUpdated() {
	r.out.Print("No updated")
}
