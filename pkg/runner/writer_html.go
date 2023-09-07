package runner

import (
	_ "embed"
	"io"

	"github.com/linuxsuren/api-testing/pkg/apispec"
	"github.com/linuxsuren/api-testing/pkg/render"
)

type htmlResultWriter struct {
	writer       io.Writer
	apiConverage apispec.APIConverage
	htmlReport   htmlReport
}

type htmlReport struct {
	Result          []ReportResult
	ResponseTimeMap map[string][]int64
}

// NewHTMLResultWriter creates a new htmlResultWriter
func NewHTMLResultWriter(writer io.Writer) ReportResultWriter {
	return &htmlResultWriter{writer: writer, htmlReport: htmlReport{
		ResponseTimeMap: map[string][]int64{},
	}}
}

// Output writes the HTML base report to target writer
func (w *htmlResultWriter) Output(result []ReportResult) (err error) {
	w.htmlReport.Result = result
	return render.RenderThenPrint("html-report", htmlReportTpl, w.htmlReport, w.writer)
}

// WithAPIConverage sets the api coverage
func (w *htmlResultWriter) WithAPIConverage(apiConverage apispec.APIConverage) ReportResultWriter {
	w.apiConverage = apiConverage
	return w
}

func (w *htmlResultWriter) WithAllRecords(records []*ReportRecord) ReportResultWriter {
	for _, record := range records {
		key := record.Name
		if _, ok := w.htmlReport.ResponseTimeMap[key]; !ok {
			w.htmlReport.ResponseTimeMap[key] = []int64{record.EndTime.Sub(record.BeginTime).Milliseconds()}
		} else {
			w.htmlReport.ResponseTimeMap[key] = append(w.htmlReport.ResponseTimeMap[key], record.EndTime.Sub(record.BeginTime).Milliseconds())
		}
	}
	return w
}

//go:embed data/html.html
var htmlReportTpl string
