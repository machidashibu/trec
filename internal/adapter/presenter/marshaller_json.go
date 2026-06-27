package presenter

import (
	"encoding/json"
)

type JsonMarshaller struct {
	indent string
}

func NewJsonMarshaller(indent string) Marshaller {
	return &JsonMarshaller{
		indent: indent,
	}
}

func (j JsonMarshaller) Marshal(_ cols, rows table) ([]byte, error) {
	if j.indent == "" {
		return json.Marshal(&rows)
	} else {
		return json.MarshalIndent(&rows, "", j.indent)
	}
}
