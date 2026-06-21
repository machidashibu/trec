package controller

import (
	"slices"
	"strconv"
	"strings"
	"trec/internal/adapter/model"
	"trec/internal/adapter/repository"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

func ParseEditOptions(args []string, config *repository.Config) (domain.RecordId, *model.EditOptions, error) {
	if len(args) < 1 {
		return 0, nil, domain.ErrorInvalidConfig
	}

	// get record id
	id, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return 0, nil, logger.Error("ParseDeleteOptions", "parse uint error", err, "val", args[0])
	}

	// get new name
	name, _ := getValueWithLabel(args, "name")
	// get new result
	result, _ := getValueWithLabel(args, "result")

	return domain.RecordId(id), model.NewEditOptions(name, result), nil

}

func getValueWithLabel(args []string, label string) (string, bool) {
	prefix := label + "="
	index := slices.IndexFunc(args, func(e string) bool { return strings.HasPrefix(e, prefix) })
	if index < 0 {
		return "", false
	}
	return strings.TrimPrefix(args[index], prefix), true
}
