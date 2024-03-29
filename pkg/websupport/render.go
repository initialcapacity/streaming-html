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
		ParseFS(resources, "resources/templates/template.gohtml", fmt.Sprintf("resources/templates/%s", fileName)))

	return temp.Execute(writer, data)
}
