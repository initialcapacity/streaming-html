package websupport

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
)

func Render(writer io.Writer, resources fs.FS, templateName string, data any) {
	_ = template.Must(template.New(fmt.Sprintf("%s.gohtml", templateName)).ParseFS(
		resources,
		"resources/templates/template.gohtml",
		fmt.Sprintf("resources/templates/%s.gohtml", templateName),
	)).Execute(writer, data)
}

func Flush(writer http.ResponseWriter) {
	writer.(http.Flusher).Flush()
}
