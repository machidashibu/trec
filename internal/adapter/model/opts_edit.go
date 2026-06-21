package model

type EditOptions struct {
	name   string
	result string
}

func NewEditOptions(name, result string) *EditOptions {
	return &EditOptions{
		name:   name,
		result: result,
	}
}

func (e EditOptions) NewName() (string, bool) {
	return e.name, e.name != ""
}

func (e EditOptions) NewResult() (string, bool) {
	return e.result, e.result != ""
}
