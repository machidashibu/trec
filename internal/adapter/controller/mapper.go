package controller

import (
	"strings"
	"trec/internal/adapter/model"
	"trec/internal/domain"
)

func MapToMode(mode string) model.Mode {
	switch strings.ToLower(mode) {
	case "r", "recording", "-r", "--recording":
		return model.ModeRecording
	case "l", "lookup", "-l", "--lookup":
		return model.ModeLookup
	case "e", "edit", "-e", "--edit":
		return model.ModeEdit
	case "d", "delete", "-d", "--delete":
		return model.ModeDelete
	case "h", "help", "-h", "--help":
		return model.ModeHelp
	default:
		return model.ModeUnknown
	}
}

func MapToOrderBy(by string) domain.OrderBy {
	switch strings.ToLower(by) {
	case "--name":
		return domain.OrderByName
	case "--start":
		return domain.OrderByStart
	case "--end":
		return domain.OrderByEnd
	case "--result":
		return domain.OrderByResult
	case "--duration":
		return domain.OrderByDuration
	case "--asc":
		return domain.OrderByAsc
	case "--desc":
		return domain.OrderByDesc
	default:
		return ""
	}
}
