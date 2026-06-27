package presenter

import (
	"fmt"
	"html"
	"strings"
)

type MarkdownMarshaller struct{}

func NewMarkdownMarshaller() Marshaller {
	return &MarkdownMarshaller{}
}

func (m MarkdownMarshaller) Marshal(cols cols, rows table) ([]byte, error) {
	var data strings.Builder

	// make header
	data.WriteString("| " + strings.Join(cols, " | ") + " |\n")
	data.WriteString("| " + strings.Repeat("--- | ", len(cols)-1) + "--- |\n")
	// make body
	for _, row := range rows {
		values := make([]string, len(cols))
		for index, col := range cols {
			if v, ok := row[col]; ok {
				values[index] = html.EscapeString(fmt.Sprintf("%v", v))
			}
		}
		data.WriteString("| " + strings.Join(values, " | ") + " |\n")
	}

	return []byte(data.String()), nil
}
