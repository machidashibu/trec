package model

import "trec/internal/domain"

type LookupOptions struct {
	format     string
	timeFormat string
	filter     domain.Filter
}

func NewLookupOptions(format, timeFormat string, filter domain.Filter) *LookupOptions {
	return &LookupOptions{
		format:     format,
		timeFormat: timeFormat,
		filter:     filter,
	}
}

func (lo LookupOptions) Format() string {
	return lo.format
}

func (lo LookupOptions) TimeFormat() string {
	return lo.timeFormat
}

func (lo LookupOptions) Filter() domain.Filter {
	return lo.filter
}
