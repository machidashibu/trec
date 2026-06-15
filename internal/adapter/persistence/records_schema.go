package persistence

import (
	"time"
	"trec/internal/adapter/model"
	"trec/internal/domain"
)

type RecordsSchema struct {
	ID        uint64    `gorm:"primaryKey,autoIncrement,not null,unique"`
	Label     string    `gorm:"not null"`
	StartTime time.Time `gorm:""`
	EndTime   time.Time `gorm:""`
	Note      string    `gorm:"not null"`
}

func (s RecordsSchema) TableName() string {
	return "records"
}

func newRecord(label string, startTime, endTime time.Time, note string) *RecordsSchema {
	return &RecordsSchema{
		Label:     label,
		StartTime: startTime,
		EndTime:   endTime,
		Note:      note,
	}
}

func fromDomain(record domain.Record) *RecordsSchema {
	return &RecordsSchema{
		Label:     record.Label(),
		StartTime: record.StartTime(),
		EndTime:   record.EndTime(),
		Note:      record.Note(),
	}
}

func toRecordList(records []RecordsSchema) domain.RecordList {
	list := model.NewRecordList()
	for _, record := range records {
		list.Add(record.ID, model.NewRecord(
			record.Label,
			record.StartTime,
			record.EndTime,
			record.Note,
		))
	}
	return list
}
