package controller

import (
	"slices"
	"trec/internal/adapter/model"
	"trec/internal/adapter/repository"
	"trec/internal/domain"
)

func ParseLookupOptions(args []string, config *repository.LookupConfig) (*model.LookupOptions, error) {
	// get style
	style := domain.LookupSimple // default:simple
	if config.DefaultFormat != "" {
		style = config.DefaultStyle
	}
	if slices.Contains(args, "--full") {
		style = domain.LookupFull
	} else if slices.Contains(args, "--collapse") {
		style = domain.LookupCollapse
	} else if slices.Contains(args, "--simple") {
		style = domain.LookupSimple
	}

	// get format
	format := domain.TableText // default;text
	if config.DefaultFormat != "" {
		format = config.DefaultFormat
	}
	if slices.Contains(args, "-c") || slices.Contains(args, "--csv") {
		format = domain.TableCsv
	} else if slices.Contains(args, "-j") || slices.Contains(args, "--json") {
		format = domain.TableJson
	} else if slices.Contains(args, "-m") || slices.Contains(args, "--markdown") {
		format = domain.TableMarkdown
	} else if slices.Contains(args, "-p") || slices.Contains(args, "--pjson") {
		format = domain.TablePrettyJson
	} else if slices.Contains(args, "-t") || slices.Contains(args, "--text") {
		format = domain.TableText
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

	return model.NewLookupOptions(style, format, timeFormat, filter), nil
}
