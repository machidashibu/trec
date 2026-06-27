package presenter

import (
	"fmt"
	"strings"
)

type TextMarshaller struct{}

func NewTextMarshaller() Marshaller {
	return &TextMarshaller{}
}

func (t *TextMarshaller) Marshal(cols cols, rows table) ([]byte, error) {
	var data strings.Builder

	// make body
	for _, row := range rows {
		values := make([]string, len(cols))
		for index, col := range cols {
			if v, ok := row[col]; ok {
				values[index] = fmt.Sprintf("%v", v)
			}
		}
		data.WriteString(strings.Join(values, "\t") + "\n")
	}

	return []byte(data.String()), nil
}
