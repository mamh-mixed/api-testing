package runner

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"

	"github.com/linuxsuren/api-testing/pkg/apispec"
)

type jsonResultWriter struct {
	writer io.Writer
}

// NewJSONResultWriter creates a new jsonResultWriter
func NewJSONResultWriter(writer io.Writer) ReportResultWriter {
	return &jsonResultWriter{writer: writer}
}

// Output writes the JSON base report to target writer
func (w *jsonResultWriter) Output(result []ReportResult) (err error) {
	jsonData, err := json.Marshal(result)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w.writer, string(jsonData))
	return
}

// WithAPIConverage sets the api coverage
func (w *jsonResultWriter) WithAPIConverage(apiConverage apispec.APIConverage) ReportResultWriter {
	return w
}

func (w *jsonResultWriter) WithAllRecords(records []*ReportRecord) ReportResultWriter {
	return w
}
