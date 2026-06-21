package manual

import (
	_ "embed"
	"trec/internal/adapter/model"
)

//go:embed general.txt
var manualGeneral string

//go:embed recording.txt
var manualRecording string

//go:embed lookup.txt
var manualLookup string

//go:embed edit.txt
var manualEdit string

//go:embed delete.txt
var manualDelete string

//go:embed help.txt
var manualHelp string

type manualOutput interface {
	Print(text string)
}

type Manual struct {
	manual map[model.Mode]string
	out    manualOutput
}

func NewManual(out manualOutput) *Manual {
	return &Manual{
		manual: map[model.Mode]string{
			model.ModeRecording: manualRecording,
			model.ModeLookup:    manualLookup,
			model.ModeEdit:      manualEdit,
			model.ModeDelete:    manualDelete,
			model.ModeHelp:      manualHelp,
		},
		out: out,
	}
}

func (m Manual) Show(mode model.Mode) int {
	if text, ok := m.manual[mode]; ok {
		m.out.Print(text)
	} else {
		m.out.Print(manualGeneral)
	}

	return 2
}
