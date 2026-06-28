package usecase

import (
	"context"
	"log/slog"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type lookupSimpleRepository interface {
	GetAll(filter domain.Filter, order []domain.Order) (domain.TestList, error)
}

type lookupSimpleReporter interface {
	Report(list domain.TestList)
}

type lookupSimpleOptions interface {
	Filter() domain.Filter
	Order() []domain.Order
}

type LookupSimple struct {
	repo     lookupSimpleRepository
	reporter lookupSimpleReporter
}

func NewLookupSimple(repo lookupSimpleRepository, reporter lookupSimpleReporter) *LookupSimple {
	return &LookupSimple{
		repo:     repo,
		reporter: reporter,
	}
}

func (uc *LookupSimple) Lookup(_ context.Context, opts lookupSimpleOptions) error {
	slog.Debug("Execute simple lookup")

	// get all items
	list, err := uc.repo.GetAll(opts.Filter(), opts.Order())
	if err != nil {
		return logger.Error("LookupSimple", "get all error", err, "filter", opts.Filter())
	}
	slog.Debug("Get all records", "len", list.Count())

	// show list
	uc.reporter.Report(list)

	slog.Debug("Finished simple lookup")
	return nil
}
