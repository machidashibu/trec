package controller

import (
	"regexp"
	"trec/internal/adapter/model"
	"trec/internal/core"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

func ParseRecordingOptions(args []string, config *core.RecordingConfig) (string, *model.RecordingOptions, error) {
	if len(args) != 1 {
		return "", nil, domain.ErrorInvalidConfig
	}

	testname := config.DefaultTestname
	if len(args) != 0 {
		testname = args[0]
	}
	if config.ValidationPattern != "" {
		matched, err := regexp.MatchString(config.ValidationPattern, testname)
		if err != nil {
			return "", nil, logger.Error("Config", "testname validattion pattern error", err, "pattern", config.ValidationPattern, "testname", testname)
		}
		if !matched {
			return "", nil, logger.Error("config", "testname validation error", domain.ErrorInvalidTestNamePattern, "pattern", config.ValidationPattern, "testname", testname)
		}
	}
	return testname, model.NewRecordingOptions(config.DefaultTimeformat), nil
}
