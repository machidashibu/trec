package presenter_test

import (
	_ "embed"
	"testing"
	"time"
	"trec/internal/adapter/presenter"

	"github.com/stretchr/testify/require"
)

func TestDurationFormatter(t *testing.T) {
	type testcase struct {
		name   string
		format string
		data   time.Duration
		result string
	}
	testcases := []testcase{
		{
			name:   "default format",
			format: "",
			data:   1*time.Hour + 23*time.Minute + 45*time.Second,
			result: "1h23m45s",
		},
		{
			name:   "#h#m#s format",
			format: "#h#m#s",
			data:   1*time.Hour + 23*time.Minute + 45*time.Second,
			result: "1h23m45s",
		},
		{
			name:   "#:#:# format",
			format: "#:#:#",
			data:   1*time.Hour + 23*time.Minute + 45*time.Second,
			result: "1:23:45",
		},
		{
			name:   "#m#s format",
			format: "#m#s",
			data:   1*time.Hour + 23*time.Minute + 45*time.Second,
			result: "83m45s",
		},
		{
			name:   "#:# format",
			format: "#:#",
			data:   1*time.Hour + 23*time.Minute + 45*time.Second,
			result: "83:45",
		},
		{
			name:   "#'#'' format",
			format: "#'#''",
			data:   1*time.Hour + 23*time.Minute + 45*time.Second,
			result: "83'45''",
		},
		{
			name:   "default format less that a hour",
			format: "",
			data:   12*time.Minute + 34*time.Second,
			result: "12m34s",
		},
		{
			name:   "#h#m#s format less that a hour",
			format: "#h#m#s",
			data:   12*time.Minute + 34*time.Second,
			result: "0h12m34s",
		},
		{
			name:   "#:#:# format less that a hour",
			format: "#:#:#",
			data:   12*time.Minute + 34*time.Second,
			result: "0:12:34",
		},
		{
			name:   "#m#s format less that a hour",
			format: "#m#s",
			data:   12*time.Minute + 34*time.Second,
			result: "12m34s",
		},
		{
			name:   "#:# format less that a hour",
			format: "#:#",
			data:   12*time.Minute + 34*time.Second,
			result: "12:34",
		},
		{
			name:   "#'#'' format less that a hour",
			format: "#'#''",
			data:   12*time.Minute + 34*time.Second,
			result: "12'34''",
		},
		{
			name:   "default format less that a minute",
			format: "",
			data:   12 * time.Second,
			result: "12s",
		},
		{
			name:   "#s format less that a minute",
			format: "#s",
			data:   12 * time.Second,
			result: "12s",
		},
		{
			name:   "#'' format less that a minute",
			format: "#''",
			data:   12 * time.Second,
			result: "12''",
		},
	}

	// testing
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.result, presenter.NewDurationFormatter(test.format).String(test.data))
		})
	}
}

//go:embed testdata/marshaller_json_without_indent.json
var marshaller_json_without_indent string

//go:embed testdata/marshaller_json_with_indent.json
var marshaller_json_with_indent string

//go:embed testdata/marshaller_markdown.md
var marshaller_markdown string

//go:embed testdata/marshaller_csv.csv
var marshaller_csv string

//go:embed testdata/marshaller_text.txt
var marshaller_text string

func TestMarshaller(t *testing.T) {
	type testcase struct {
		name   string
		target presenter.Marshaller
		result string
	}
	testcases := []testcase{
		{
			name:   "json without indent",
			target: presenter.NewJsonMarshaller(""),
			result: marshaller_json_without_indent,
		},
		{
			name:   "json with indent",
			target: presenter.NewJsonMarshaller("  "),
			result: marshaller_json_with_indent,
		},
		{
			name:   "markdown",
			target: presenter.NewMarkdownMarshaller(),
			result: marshaller_markdown,
		},
		{
			name:   "csv",
			target: presenter.NewCsvMarshaller(),
			result: marshaller_csv,
		},
		{
			name:   "text",
			target: presenter.NewTextMarshaller(),
			result: marshaller_text,
		},
	}

	// testdata
	var (
		cols = []string{"col1", "col2", "col3"}
		rows = []map[string]any{
			{"col1": "val1", "col2": 1, "col3": time.Date(2026, 1, 23, 1, 23, 45, 0, time.Local)},
			{"col1": "val 2", "col2": 2.0, "col3": time.Duration(1234)},
			{"col1": `val "3"`, "col2": -3.0, "col3": nil},
			{"col1": "val4"},
		}
	)

	// testing
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.target.Marshal(cols, rows)
			require.NoError(t, err)
			// t.Log(string(result))
			require.Equal(t, test.result, string(result))
		})
	}
}
