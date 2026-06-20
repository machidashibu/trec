package model

type Mode string

const (
	ModeRecording Mode = "recording"
	ModeLookup    Mode = "lookup"
	ModeDelete    Mode = "delete"
	ModeHelp      Mode = "help"
	ModeUnknown   Mode = "unknown"
)
