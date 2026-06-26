package model

import "trec/internal/domain"

type testListItem struct {
	id   domain.RecordId
	test domain.Test
}

type TestList struct {
	items []testListItem
}

func NewRecordList() *TestList {
	return &TestList{
		items: []testListItem{},
	}
}

func (tl TestList) Count() int {
	return len(tl.items)
}

func (tl TestList) Get(index int) (domain.Test, domain.RecordId) {
	return tl.items[index].test, tl.items[index].id
}

func (tl *TestList) Add(id domain.RecordId, test domain.Test) {
	tl.items = append(tl.items, testListItem{
		id:   id,
		test: test,
	})
}
