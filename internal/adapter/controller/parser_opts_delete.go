package controller

import (
	"strconv"
	"trec/internal/core"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

func ParseDeleteOptions(args []string, config *core.Config) (domain.RecordId, error) {
	if len(args) != 1 {
		return 0, domain.ErrorInvalidConfig
	}

	// get record id
	if id, err := strconv.ParseUint(args[0], 10, 64); err != nil {
		return 0, logger.Error("ParseDeleteOptions", "parse uint error", err, "val", args[0])
	} else {
		return domain.RecordId(id), nil
	}
}
