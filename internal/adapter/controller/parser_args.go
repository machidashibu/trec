package controller

import "trec/internal/core"

func ParseArgs(args []string) (core.Mode, []string) {
	if len(args) == 0 {
		return core.ModeRecording, args // default: recoding mode
	}

	// parse mode
	switch args[0] {
	case "-l", "--lookup":
		return core.ModeLookup, args[1:]
	case "-d", "--delete":
		return core.ModeDelete, args[1:]
	case "-h", "--help":
		return core.ModeHelp, args[1:]
	// case "-r", "--recording":
	default: // default: recoding mode
		if args[0] == "-r" || args[0] == "--recording" {
			// specified mode
			return core.ModeRecording, args[1:]
		} else {
			// not specified mode
			return core.ModeRecording, args
		}
	}
}
