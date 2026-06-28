package model

import "trec/internal/domain"

type LookupOptions struct {
	style      domain.LookupStyle
	format     domain.TableFormat
	timeFormat string
	filter     domain.Filter
	order      domain.Order
}

func NewLookupOptions(style domain.LookupStyle, format domain.TableFormat, timeFormat string, filter domain.Filter, order domain.Order) *LookupOptions {
	return &LookupOptions{
		style:      style,
		format:     format,
		timeFormat: timeFormat,
		filter:     filter,
		order:      order,
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

func (l LookupOptions) Order() []domain.Order {
	if l.order == nil {
		return []domain.Order{} // no order
	} else {
		return NewOrderList(l.order)
	}
}
