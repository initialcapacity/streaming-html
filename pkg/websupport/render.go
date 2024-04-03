package websupport

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

func Render(writer http.ResponseWriter, resources fs.FS, templateName string, data any) error {
	fileName := fmt.Sprintf("%s.gohtml", templateName)

	return template.Must(template.New(fileName).
		ParseFS(resources, templatePath("template.gohtml"), templatePath(fileName))).
		Execute(writer, data)
}

func templatePath(fileName string) string {
	return fmt.Sprintf("resources/templates/%s", fileName)
}
