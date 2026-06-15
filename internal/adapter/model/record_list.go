package model

import "trec/internal/domain"

type recordListItem struct {
	id     uint64
	record domain.Record
}

type RecordList struct {
	items []recordListItem
}

func NewRecordList() *RecordList {
	return &RecordList{
		items: []recordListItem{},
	}
}

func (rl RecordList) Count() int {
	return len(rl.items)
}

func (rl RecordList) Get(index int) (uint64, domain.Record) {
	return rl.items[index].id, rl.items[index].record
}

func (rl *RecordList) Add(id uint64, record domain.Record) {
	rl.items = append(rl.items, recordListItem{
		id:     id,
		record: record,
	})
}
