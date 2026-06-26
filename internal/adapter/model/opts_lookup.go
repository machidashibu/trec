package model

import "trec/internal/domain"

type LookupOptions struct {
	format     domain.LookupFormat
	timeFormat string
	filter     domain.Filter
}

func NewLookupOptions(format domain.LookupFormat, timeFormat string, filter domain.Filter) *LookupOptions {
	return &LookupOptions{
		format:     format,
		timeFormat: timeFormat,
		filter:     filter,
	}
}

func (lo LookupOptions) Format() domain.LookupFormat {
	return lo.format
}

func (lo LookupOptions) TimeFormat() string {
	return lo.timeFormat
}

func (lo LookupOptions) Filter() domain.Filter {
	return lo.filter
}
