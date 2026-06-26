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

	// get filter.today
	filter := model.NewNoFilter() // default:no filter
	if !slices.Contains(args, "--no-filter") {
		// get today
		today := config.DefaultFilter.StartTimeToday
		if slices.Contains(args, "--all-days") {
			today = false
		} else if slices.Contains(args, "--today") {
			today = true
		}

		// get latest only
		latestOnly := config.DefaultFilter.LatestOnlyPerTestname
		if slices.Contains(args, "--latest-only") {
			latestOnly = true
		}

		filter = model.NewFilter(today, latestOnly)
	}
	return model.NewLookupOptions(format, timeFormat, filter), nil
}
