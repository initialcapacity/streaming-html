package websupport

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

func Render(writer http.ResponseWriter, resources fs.FS, templateName string, data any) error {
	fileName := fmt.Sprintf("%s.gohtml", templateName)
	functions := template.FuncMap{
		"flush": func() (string, error) {
			err := flush(writer)
			return "", err
		},
	}
	temp := template.Must(template.New(fileName).
		Funcs(functions).
		ParseFS(
			resources,
			"resources/templates/template.gohtml",
			fmt.Sprintf("resources/templates/%s", fileName),
		))

	err := temp.Execute(writer, data)
	if err != nil {
		return err
	}

	return flush(writer)
}

func flush(writer http.ResponseWriter) error {
	return http.NewResponseController(writer).Flush()
}
