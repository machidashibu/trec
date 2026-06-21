package repository

import (
	"time"
	"trec/internal/adapter/model"
	"trec/internal/domain"
)

type TestResultSchema struct {
	ID        uint64    `gorm:"primaryKey,autoIncrement,not null,unique"`
	Name      string    `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	Result    string    `gorm:"not null"`
}

func (s TestResultSchema) TableName() string {
	return "test_result"
}

func newRecord(name string, startTime, endTime time.Time, result string) *TestResultSchema {
	return &TestResultSchema{
		Name:      name,
		StartTime: startTime,
		EndTime:   endTime,
		Result:    result,
	}
}

func fromDomain(record domain.Test) *TestResultSchema {
	return &TestResultSchema{
		Name:      record.Name(),
		StartTime: record.StartTime(),
		EndTime:   record.EndTime(),
		Result:    record.Result(),
	}
}

func toTestList(records []TestResultSchema) domain.TestList {
	list := model.NewRecordList()
	for _, record := range records {
		list.Add(domain.RecordId(record.ID), model.NewTest(
			record.Name,
			record.StartTime,
			record.EndTime,
			record.Result,
		))
	}
	return list
}
