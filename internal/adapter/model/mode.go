package model

type Mode string

const (
	ModeRecording Mode = "recording"
	ModeLookup    Mode = "lookup"
	ModeEdit      Mode = "edit"
	ModeDelete    Mode = "delete"
	ModeHelp      Mode = "help"
	ModeUnknown   Mode = "unknown"
)
