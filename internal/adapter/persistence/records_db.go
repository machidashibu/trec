package persistence

import (
	"log/slog"
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

func (d *RecordsDatabase) GetAll(filter domain.Filter) (domain.RecordList, error) {
	slog.Debug("Called RecordsDatabase.GetAll", "filter", filter)

	db := d.infra.DB()
	// set order
	// if strings.HasPrefix(string(order), string(domain.OrderByDuration)) {
	// 	db = db.Order("julianday(end_time) - julianday(start_time) ASC")
	// } else {
	// 	db = db.Order(order)
	// }

	// set filter
	if filter.Today() {
		// TODO: DB dependent
		slog.Debug("filtered by start_id at today")
		today := time.Now().Format("2006-01-02")
		db = db.Where("DATE(start_time) = ?", today)
	}
	if filter.LatestOnly() {
		// keep only the latest record for each unique label
		sub := d.infra.DB().Model(&RecordsSchema{}).
			Select("label, MAX(start_time) AS max_start_time").
			Group("label")

		db = db.Joins(
			"JOIN (?) AS latest ON records.label = latest.label AND records.start_time = latest.max_start_time",
			sub,
		)
	}

	// get records
	var records []RecordsSchema
	if err := db.Find(&records).Error; err != nil {
		return nil, logger.Error("RecordsDatabase", "find error", err)
	}
	slog.Debug("Get all records", "len", len(records))

	return toDomainArray(records), nil
}
