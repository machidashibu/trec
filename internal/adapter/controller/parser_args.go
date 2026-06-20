package controller

import "trec/internal/adapter/model"

func ParseArgs(args []string) (model.Mode, []string) {
	if len(args) == 0 {
		return model.ModeRecording, args // default: recoding mode
	}

	// parse mode
	switch args[0] {
	case "-l", "--lookup":
		return model.ModeLookup, args[1:]
	case "-d", "--delete":
		return model.ModeDelete, args[1:]
	case "-h", "--help":
		return model.ModeHelp, args[1:]
	// case "-r", "--recording":
	default: // default: recoding mode
		if args[0] == "-r" || args[0] == "--recording" {
			// specified mode
			return model.ModeRecording, args[1:]
		} else {
			// not specified mode
			return model.ModeRecording, args
		}
	}
}
