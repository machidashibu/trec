package usecase

import (
	"context"
	"log/slog"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type lookupFullRepository interface {
	GetAll(filter domain.Filter, order []domain.Order) (domain.TestList, error)
}

type lookupFullReporter interface {
	Report(list domain.TestList)
}

type lookupFullOptions interface {
	Filter() domain.Filter
	Order() []domain.Order
}

type LookupFull struct {
	repo     lookupFullRepository
	reporter lookupFullReporter
}

func NewLookupFull(repo lookupFullRepository, reporter lookupFullReporter) *LookupFull {
	return &LookupFull{
		repo:     repo,
		reporter: reporter,
	}
}

func (uc *LookupFull) Lookup(_ context.Context, opts lookupFullOptions) error {
	slog.Debug("Execute full lookup")

	// get all items
	list, err := uc.repo.GetAll(opts.Filter(), opts.Order())
	if err != nil {
		return logger.Error("LookupFull", "get all error", err, "filter", opts.Filter())
	}
	slog.Debug("Get all records", "len", list.Count())

	// show list
	uc.reporter.Report(list)

	slog.Debug("Finished full lookup")
	return nil
}
