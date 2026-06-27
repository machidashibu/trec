package model

import "trec/internal/domain"

type LookupOptions struct {
	style      domain.LookupStyle
	format     domain.TableFormat
	timeFormat string
	filter     domain.Filter
}

func NewLookupOptions(style domain.LookupStyle, format domain.TableFormat, timeFormat string, filter domain.Filter) *LookupOptions {
	return &LookupOptions{
		style:      style,
		format:     format,
		timeFormat: timeFormat,
		filter:     filter,
	}
}

func (lo LookupOptions) Style() domain.LookupStyle {
	return lo.style
}

func (lo LookupOptions) Format() domain.TableFormat {
	return lo.format
}

func (lo LookupOptions) TimeFormat() string {
	return lo.timeFormat
}

func (lo LookupOptions) Filter() domain.Filter {
	return lo.filter
}
