package repository

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

type TestResultDatabase struct {
	infra infraDB
}

// NOE: It can not create an instance from external. (Should be created by factory)
func newTestResultDatabase(infra infraDB) *TestResultDatabase {
	return &TestResultDatabase{infra: infra}
}

func (d *TestResultDatabase) ensureTable() error {
	if err := d.infra.DB().AutoMigrate(&TestResultSchema{}); err != nil {
		return logger.Error("RecordsDatabase", "auto-migrate error", err)
	}
	return nil
}

func (d *TestResultDatabase) Add(name string, start time.Time, end time.Time, result string) (domain.Test, error) {
	slog.Debug("Called RecordsDatabase.Add", "name", name, "start", start, "end", end, "result", result)
	record := newRecord(name, start, end, result)

	if err := d.infra.DB().Create(record).Error; err != nil {
		return nil, logger.Error("RecordsDatabase", "create error", err, "record", record)
	}
	slog.Debug("Created record", "record", record)

	return model.NewRecord(record.Name, record.StartTime, record.EndTime, record.Result), nil
}

func (d *TestResultDatabase) GetAll(filter domain.Filter) (domain.TestList, error) {
	slog.Debug("Called TestResultDatabase.GetAll", "filter", filter)

	db := d.infra.DB()
	// set filter
	if filter.Today() {
		now := time.Now()
		startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		endOfDay := startOfDay.Add(24 * time.Hour)
		db = db.Where("start_time >= ? AND start_time < ?", startOfDay, endOfDay)
	}
	if filter.LatestOnly() {
		// keep only the latest record for each unique name
		sub := d.infra.DB().Model(&TestResultSchema{}).
			Select("name, MAX(start_time) AS max_start_time").
			Group("name")

		db = db.Joins(
			"JOIN (?) AS latest ON test_result.name = latest.name AND test_result.start_time = latest.max_start_time",
			sub,
		)
	}

	// set order
	// if strings.HasPrefix(string(order), string(domain.OrderByDuration)) {
	// 	db = db.Order("julianday(end_time) - julianday(start_time) ASC")
	// } else {
	// 	db = db.Order(order)
	// }

	// get records
	var records []TestResultSchema
	if err := db.Find(&records).Error; err != nil {
		return nil, logger.Error("RecordsDatabase", "find error", err)
	}
	slog.Debug("Get all records", "len", len(records))

	return toRecordList(records), nil
}

func (db *TestResultDatabase) Delete(id domain.RecordId) error {
	slog.Debug("Called TestResultDatabase.Delete", "id", id)

	return db.infra.DB().Where("id = ?", id).Delete(TestResultSchema{}).Error
}
