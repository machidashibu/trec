package presenter

import (
	"io"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type TableReporter struct {
	w    io.Writer
	m    Marshaller
	cols []string
	rows []map[string]any
}

func NewTableReporter(w io.Writer, m Marshaller) *TableReporter {
	return &TableReporter{
		w: w,
		m: m,
	}
}

func (t *TableReporter) Header(cols ...string) error {
	t.cols = cols
	t.rows = []map[string]any{}
	return nil
}

func (t *TableReporter) Row(values ...any) error {
	if t.cols == nil || t.rows == nil {
		return domain.ErrorInvalidOperation
	}

	row := map[string]any{}
	for index, value := range values {
		row[t.cols[index]] = value
	}

	t.rows = append(t.rows, row)

	return nil
}

func (t *TableReporter) Save() error {
	data, err := t.m.Marshal(t.cols, t.rows)
	if err != nil {
		return logger.Error("TableReporter", "marshal error", err, "cols", t.cols, "rows", t.rows)
	}

	_, err = t.w.Write(data)
	if err != nil {
		return logger.Error("TableReporter", "write error", err)
	}
	return nil
}
