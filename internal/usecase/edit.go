package usecase

import (
	"log/slog"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type editErepository interface {
	EditName(id domain.RecordId, name string) error
	EditResult(id domain.RecordId, result string) error
	Edit(id domain.RecordId, name, result string) error
}

type editReporter interface {
	ReportUpdatedName(name string)
	ReportUpdatedResult(result string)
	ReportNoUpdated()
}

type editOptions interface {
	NewName() (string, bool)
	NewResult() (string, bool)
}

type Edit struct {
	repo     editErepository
	reporter editReporter
}

func NewEdit(repo editErepository, reporter editReporter) *Edit {
	return &Edit{
		repo:     repo,
		reporter: reporter,
	}
}

func (uc Edit) Edit(id domain.RecordId, opts editOptions) error {
	slog.Debug("Execute edit", "id", id, "opts", opts)

	name, existName := opts.NewName()
	result, existResult := opts.NewResult()
	if existName && existResult {
		if err := uc.repo.Edit(id, name, result); err != nil {
			return logger.Error("Edit", "edit error", err, "id", id, "name", name, "result", result)
		}
		uc.reporter.ReportUpdatedName(name)
		uc.reporter.ReportUpdatedResult(result)
		slog.Debug("Edited", "id", id, "name", name, "result", result)
	} else if existName {
		if err := uc.repo.EditName(id, name); err != nil {
			return logger.Error("Edit", "edit name error", err, "id", id, "name", name)
		}
		uc.reporter.ReportUpdatedName(name)
		slog.Debug("Edited name", "id", id, "name", name)
	} else if existResult {
		if err := uc.repo.EditResult(id, result); err != nil {
			return logger.Error("Edit", "edit result error", err, "id", id, "result", result)
		}
		uc.reporter.ReportUpdatedResult(result)
		slog.Debug("Edited result", "id", id, "result", result)
	} else {
		uc.reporter.ReportNoUpdated()
		slog.Debug("no update")
	}

	slog.Debug("Finished edit")
	return nil
}
