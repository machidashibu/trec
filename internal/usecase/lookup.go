package usecase

import (
	"log/slog"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type lookupRepository interface {
	GetAll(order domain.OrderBy) (domain.RecordList, error)
}

type lookupReporter interface {
	Report(list domain.RecordList)
}

type lookupOptions interface {
	LookupOrder() domain.OrderBy
}

type Lookup struct {
	repo     lookupRepository
	reporter lookupReporter
}

func NewLookup(repo lookupRepository, reporter lookupReporter) *Lookup {
	return &Lookup{
		repo:     repo,
		reporter: reporter,
	}
}

func (uc *Lookup) Lookup(opts lookupOptions) error {
	slog.Debug("Execute lookup")

	// get all items
	list, err := uc.repo.GetAll(opts.LookupOrder())
	if err != nil {
		return logger.Error("Lookup", "get all error", err)
	}
	slog.Debug("Get all records", "len", len(list))

	// show list
	uc.reporter.Report(list)

	slog.Debug("Finished lookup")
	return nil
}
