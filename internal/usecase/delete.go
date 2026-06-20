package usecase

import (
	"context"
	"log/slog"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type deletErepository interface {
	Delete(id domain.RecordId) error
}

type deleteReporter interface {
	Report(id domain.RecordId)
}

type Delete struct {
	repo     deletErepository
	reporter deleteReporter
}

func NewDelete(repo deletErepository, reporter deleteReporter) *Delete {
	return &Delete{
		repo:     repo,
		reporter: reporter,
	}
}

func (uc *Delete) Delete(_ context.Context, id domain.RecordId) error {
	slog.Debug("Execute delete", "id", id)

	if err := uc.repo.Delete(id); err != nil {
		return logger.Error("delete", "delete error", err, "id", id)
	}
	uc.reporter.Report(id)
	slog.Debug("Deleted record", "id", id)

	slog.Debug("Finished delete", "id", id)
	return nil
}
