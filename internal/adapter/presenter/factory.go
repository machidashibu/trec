package presenter

import (
	"os"
	"trec/internal/domain"
)

type tableReporterOptions interface {
	Format() domain.TableFormat
}

func CreateTableReporter(opts tableReporterOptions) tableReporter {
	var m Marshaller
	switch opts.Format() {
	case domain.TableCsv:
		m = NewCsvMarshaller()
	case domain.TableJson:
		m = NewJsonMarshaller("")
	case domain.TableMarkdown:
		m = NewMarkdownMarshaller()
	case domain.TablePrettyJson:
		m = NewJsonMarshaller("  ")
	// case domain.TableText:
	default:
		m = NewTextMarshaller()
	}

	return NewTableReporter(os.Stdout, m)
}
