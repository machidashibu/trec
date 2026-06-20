package controller

import (
	"slices"
	"trec/internal/adapter/model"
	"trec/internal/adapter/repository"
)

func ParseLookupOptions(args []string, config *repository.LookupConfig) (*model.LookupOptions, error) {
	// get format
	format := "simple"
	if config.DefaultFormat != "" {
		format = config.DefaultFormat
	}
	if slices.Contains(args, "--format-full") {
		format = "full"
	}

	// get time format
	timeFormat := config.DefaultTimeformat

	// get filter.today
	filter := model.NewNoFilter() // default: no filter
	if !slices.Contains(args, "--no-filter") {
		// get today
		today := config.DefaultFilter.StartTimeToday
		if slices.Contains(args, "--all-days") {
			config.DefaultFilter.StartTimeToday = false
		} else if slices.Contains(args, "--today") {
			config.DefaultFilter.StartTimeToday = true
		}

		// get latest only
		latestOnly := config.DefaultFilter.LatestOnlyPerTestname
		if slices.Contains(args, "--latest-only") {
			config.DefaultFilter.LatestOnlyPerTestname = true
		}

		filter = model.NewFilter(today, latestOnly)
	}
	return model.NewLookupOptions(format, timeFormat, filter), nil
}
