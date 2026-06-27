package presenter

type cols []string
type table []map[string]any

type Marshaller interface {
	Marshal(cols cols, rows table) ([]byte, error)
}

type tableReporter interface {
	Header(cols ...string) error
	Row(values ...any) error
	Save() error
}
