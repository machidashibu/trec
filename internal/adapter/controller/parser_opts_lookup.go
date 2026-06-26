package controller

import (
	"slices"
	"trec/internal/adapter/model"
	"trec/internal/adapter/repository"
	"trec/internal/domain"
)

func ParseLookupOptions(args []string, config *repository.LookupConfig) (*model.LookupOptions, error) {
	// get format
	format := domain.LookupSimple // default:simple
	if config.DefaultFormat != "" {
		format = config.DefaultFormat
	}
	if slices.Contains(args, "--full") {
		format = domain.LookupFull
	} else if slices.Contains(args, "--collapse") {
		format = domain.LookupCollapse
	}

	// get time format
	timeFormat := config.DefaultTimeformat

	// NOTE: Does not use configuration file options if specified command line arguments
	defaultConfig := new(repository.LookupConfig)
	if len(args) == 0 {
		defaultConfig = config
	}

	// get filter
	filter := model.NewNoFilter() // default:no filter
	if !slices.Contains(args, "--no-filter") {
		// get today
		today := defaultConfig.DefaultFilter.StartTimeToday
		if slices.Contains(args, "--all-days") {
			today = false
		} else if slices.Contains(args, "--today") {
			today = true
		}

		// get latest only
		latestOnly := defaultConfig.DefaultFilter.LatestOnlyPerTestname
		if slices.Contains(args, "--latest-only") {
			latestOnly = true
		}

		filter = model.NewFilter(today, latestOnly)
	}
	return model.NewLookupOptions(format, timeFormat, filter), nil
}
