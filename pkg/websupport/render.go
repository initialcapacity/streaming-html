package websupport

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

func Render(writer http.ResponseWriter, resources fs.FS, templateName string, data any) {
	fileName := fmt.Sprintf("%s.gohtml", templateName)
	functions := template.FuncMap{
		"flush": func() string {
			flush(writer)
			return ""
		},
	}
	temp := template.Must(template.New(fileName).
		Funcs(functions).
		ParseFS(
			resources,
			"resources/templates/template.gohtml",
			fmt.Sprintf("resources/templates/%s", fileName),
		))

	_ = temp.Execute(writer, data)
	flush(writer)
}

func flush(writer http.ResponseWriter) {
	writer.(http.Flusher).Flush()
}
