package websupport

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

func Render(writer http.ResponseWriter, resources fs.FS, templateName string, data any) error {
	fileName := fmt.Sprintf("%s.gohtml", templateName)

	writer.Header().Set("Transfer-Encoding", "chunked")
	return template.Must(template.New(fileName).
		Funcs(template.FuncMap{
			"defer": func(value any) (any, error) {
				err := flush(writer)
				return value, err
			},
		}).
		ParseFS(resources, "resources/templates/template.gohtml", fmt.Sprintf("resources/templates/%s", fileName))).
		Execute(writer, data)
}

func flush(writer http.ResponseWriter) error {
	return http.NewResponseController(writer).Flush()
}
