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

type Delete struct {
	repo deletErepository
}

func NewDelete(repo deletErepository) *Delete {
	return &Delete{
		repo: repo,
	}
}

func (uc *Delete) Delete(_ context.Context, id domain.RecordId) error {
	slog.Debug("Execute delete", "id", id)

	if err := uc.repo.Delete(id); err != nil {
		return logger.Error("delete", "delete error", err, "id", id)
	}
	slog.Debug("Delete record", "id", id)

	slog.Debug("Finished delete", "id", id)
	return nil
}
