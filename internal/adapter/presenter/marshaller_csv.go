package presenter

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type CsvMarshaller struct{}

func NewCsvMarshaller() Marshaller {
	return &CsvMarshaller{}
}

func (c CsvMarshaller) Marshal(cols cols, rows table) ([]byte, error) {
	var b strings.Builder
	w := csv.NewWriter(&b)

	if err := w.Write(cols); err != nil {
		return nil, err
	}

	for _, row := range rows {
		values := make([]string, len(cols))
		for i, col := range cols {
			values[i] = fmt.Sprintf("%v", row[col])
		}
		if err := w.Write(values); err != nil {
			return nil, err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}

	return []byte(b.String()), nil
}
