package model

type Filter struct {
	today      bool
	latestOnly bool
}

func NewFilter(today, latestOnly bool) *Filter {
	return &Filter{
		today:      today,
		latestOnly: latestOnly,
	}
}

func NewNoFilter() *Filter {
	return &Filter{}
}

func (f Filter) Today() bool {
	return f.today
}

func (f Filter) LatestOnly() bool {
	return f.latestOnly
}
