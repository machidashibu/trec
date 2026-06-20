package controller

import (
	"strings"
	"trec/internal/adapter/model"
)

func MapToMode(mode string) model.Mode {
	switch strings.ToLower(mode) {
	case "r", "recording", "-r", "--recording":
		return model.ModeRecording
	case "l", "lookup", "-l", "--lookup":
		return model.ModeLookup
	case "d", "delete", "-d", "--delete":
		return model.ModeDelete
	case "h", "help", "-h", "--help":
		return model.ModeHelp
	default:
		return model.ModeUnknown
	}
}
