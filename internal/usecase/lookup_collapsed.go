package usecase

import (
	"context"
	"log/slog"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type lookupCollapsedRepository interface {
	GetCollapsed(filter domain.Filter, order []domain.Order) (domain.CollapsedTestList, error)
}

type lookupCollapsedReporter interface {
	Report(list domain.CollapsedTestList)
}

type lookupCollapsedOptions interface {
	Filter() domain.Filter
	Order() []domain.Order
}

type LookupCollapsed struct {
	repo     lookupCollapsedRepository
	reporter lookupCollapsedReporter
}

func NewLookupCollapsed(repo lookupCollapsedRepository, reporter lookupCollapsedReporter) *LookupCollapsed {
	return &LookupCollapsed{
		repo:     repo,
		reporter: reporter,
	}
}

func (uc *LookupCollapsed) Lookup(_ context.Context, opts lookupCollapsedOptions) error {
	slog.Debug("Execute collapsed lookup")

	// get all items
	list, err := uc.repo.GetCollapsed(opts.Filter(), opts.Order())
	if err != nil {
		return logger.Error("LookupCollapsed", "get all error", err, "filter", opts.Filter())
	}
	slog.Debug("Get all records", "len", len(list))

	// show list
	uc.reporter.Report(list)

	slog.Debug("Finished collapsed lookup ")
	return nil
}
