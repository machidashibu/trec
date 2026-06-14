package persistence

import (
	"log/slog"
	"strings"
	"time"
	"trec/internal/adapter/model"
	"trec/internal/core/logger"
	"trec/internal/domain"

	"gorm.io/gorm"
)

type infraDB interface {
	DB() *gorm.DB
}

type RecordsDatabase struct {
	infra infraDB
}

// NOE: It can not create an instance from external. (Should be created by factory)
func newRecordsDatabase(infra infraDB) *RecordsDatabase {
	return &RecordsDatabase{infra: infra}
}

func (d *RecordsDatabase) ensureTable() error {
	if err := d.infra.DB().AutoMigrate(&RecordsSchema{}); err != nil {
		return logger.Error("RecordsDatabase", "auto-migrate error", err)
	}
	return nil
}

func (d *RecordsDatabase) Add(label string, start time.Time, end time.Time, memo string) (domain.Record, error) {
	slog.Debug("Called RecordsDatabase.Add", "label", label, "start", start, "end", end, "memo", memo)
	record := newRecord(label, start, end, memo)

	if err := d.infra.DB().Create(record).Error; err != nil {
		return nil, logger.Error("RecordsDatabase", "create error", err, "record", record)
	}
	slog.Debug("Created record", "record", record)

	return model.NewRecord(record.Label, record.StartTime, record.EndTime, record.Note), nil
}

func (d *RecordsDatabase) GetAll(order domain.OrderBy) (domain.RecordList, error) {
	slog.Debug("Called RecordsDatabase.GetAll", "order", order)

	db := d.infra.DB()
	// set order
	if strings.HasPrefix(string(order), string(domain.OrderByDuration)) {
		db = db.Order("julianday(end_time) - julianday(start_time) ASC")
	} else {
		db = db.Order(order)
	}

	// get records
	var records []RecordsSchema
	if err := db.Find(&records).Error; err != nil {
		return nil, logger.Error("RecordsDatabase", "find error", err)
	}
	slog.Debug("Get all records", "len", len(records))

	return toDomainArray(records), nil
}
