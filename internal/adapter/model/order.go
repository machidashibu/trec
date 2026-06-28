package model

import "trec/internal/domain"

type Order struct {
	column domain.OrderBy
	desc   bool
}

func NewOrder(column domain.OrderBy, desc bool) domain.Order {
	return &Order{
		column: column,
		desc:   desc,
	}
}

func NewOrderList(orders ...domain.Order) []domain.Order {
	return orders
}

func (o Order) Column() domain.OrderBy {
	return o.column
}

func (o Order) Desc() bool {
	return o.desc
}
