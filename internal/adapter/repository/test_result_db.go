package repository

import (
	"fmt"
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

// Add creates new record with name, start time, end time and result, return created item.
func (d *TestResultDatabase) Add(name string, start time.Time, end time.Time, result string) (domain.Test, error) {
	slog.Debug("Called RecordsDatabase.Add", "name", name, "start", start, "end", end, "result", result)

	record := newRecord(name, start, end, result)
	if err := d.infra.DB().Create(record).Error; err != nil {
		return nil, logger.Error("RecordsDatabase", "create error", err, "record", record)
	}
	slog.Debug("Created record", "record", record)

	return model.NewTest(record.Name, record.StartTime, record.EndTime, record.Result), nil
}

func (d *TestResultDatabase) applyFilter(filter domain.Filter) *gorm.DB {
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

		query := fmt.Sprintf(
			"JOIN (?) AS latest ON %s.name = latest.name AND %s.start_time = latest.max_start_time",
			(&TestResultSchema{}).TableName(),
			(&TestResultSchema{}).TableName(),
		)
		db = db.Joins(query, sub)
	}

	// set order
	// if strings.HasPrefix(string(order), string(domain.OrderByDuration)) {
	// 	db = db.Order("julianday(end_time) - julianday(start_time) ASC")
	// } else {
	// 	db = db.Order(order)
	// }

	return db
}

func (d *TestResultDatabase) GetAll(filter domain.Filter) (domain.TestList, error) {
	slog.Debug("Called TestResultDatabase.GetAll", "filter", filter)

	db := d.applyFilter(filter)

	// get records
	var records []TestResultSchema
	if err := db.Find(&records).Error; err != nil {
		return nil, logger.Error("TestResultDatabase", "find error", err)
	}
	slog.Debug("Get all records", "len", len(records))

	return toTestList(records), nil
}

func (d *TestResultDatabase) GetCollapsed(filter domain.Filter) (domain.CollapsedTestList, error) {
	slog.Debug("Called TestResultDatabase.GetCollapsed", "filter", filter)

	db := d.applyFilter(filter)
	db = db.Model(&TestResultSchema{}).
		Select("MIN(id) AS id, name, COUNT(*) AS test_count, SUM((julianday(end_time) - julianday(start_time)) * 86400.0) AS total_duration_secs")
	db = db.Group("name").Order("name ASC")

	// get records
	var records []collapsedTestResult
	if err := db.Find(&records).Error; err != nil {
		return nil, logger.Error("TestResultDatabase", "find error", err)
	}
	slog.Debug("Get collapsed records", "len", len(records))

	return toCollapsedTestList(records), nil
}

// GetById gets record by id.
func (d *TestResultDatabase) GetById(id domain.RecordId) (domain.Test, error) {
	slog.Debug("Called TestResultDatabase.GetById", "id", id)

	var record TestResultSchema
	if err := d.infra.DB().Where("id = ?", id).First(&record).Error; err != nil {
		return nil, logger.Error("TestResultDatabase", "first error", err, "id", id)
	}

	return model.NewTest(record.Name, record.StartTime, record.EndTime, record.Result), nil
}

// EditName updates columns of name with new value.
func (d *TestResultDatabase) EditName(id domain.RecordId, name string) error {
	slog.Debug("Called TestResultDatabase.EditResult", "id", id, "name", name)

	// update name
	if err := d.infra.DB().Model(&TestResultSchema{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return logger.Error("TestResultDatabase", "update error", err, "id", id, "name", name)
	}
	slog.Debug("Updated name column value")

	return nil
}

// EditResult updates columns of result with new value.
func (d *TestResultDatabase) EditResult(id domain.RecordId, result string) error {
	slog.Debug("Called TestResultDatabase.EditResult", "id", id, "result", result)

	// update result
	if err := d.infra.DB().Model(&TestResultSchema{}).Where("id = ?", id).Update("result", result).Error; err != nil {
		return logger.Error("TestResultDatabase", "update error", err, "id", id, "result", result)
	}
	slog.Debug("Updated result column value")

	return nil
}

// Edit updates columns of name and result with new value.
// It revert DB if update name or result that is failed.
func (d *TestResultDatabase) Edit(id domain.RecordId, name, result string) error {
	slog.Debug("Called TestResultDatabase.EditResult", "id", id, "name", name, "result", result)

	return d.infra.DB().Model(&TestResultSchema{}).Transaction(func(tx *gorm.DB) error {
		// update name
		if err := tx.Where("id = ?", id).Update("name", name).Error; err != nil {
			return logger.Error("TestResultDatabase", "update error", err, "id", id, "name", name)
		}
		slog.Debug("Updated name column value")
		// update result
		if err := tx.Where("id = ?", id).Update("result", result).Error; err != nil {
			return logger.Error("TestResultDatabase", "update error", err, "id", id, "result", result)
		}
		slog.Debug("Updated result column value")
		return nil
	})
}

// Delete deletes record by ID.
func (d *TestResultDatabase) Delete(id domain.RecordId) error {
	slog.Debug("Called TestResultDatabase.Delete", "id", id)

	// delete record
	if err := d.infra.DB().Where("id = ?", id).Delete(TestResultSchema{}).Error; err != nil {
		return logger.Error("TestResultDatabase", "update error", err, "id", id)
	}
	slog.Debug("Deleted record")

	return nil
}
