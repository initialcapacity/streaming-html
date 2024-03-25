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
	temp := template.Must(template.New(fileName).
		Funcs(template.FuncMap{
			"flush": func(value any) any {
				go func() {
					// flush immediately after the value is returned
					// TODO: address the data race
					_ = flush(writer)
				}()

				return value
			},
		}).
		ParseFS(resources, "resources/templates/template.gohtml", fmt.Sprintf("resources/templates/%s", fileName)))

	AddFlush(temp)

	return temp.Execute(writer, data)
}

func flush(writer http.ResponseWriter) error {
	return http.NewResponseController(writer).Flush()
}
