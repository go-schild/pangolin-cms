package cms

import (
	"html/template"
	"path"
)

var adminTemplates = map[string]*template.Template{}

func init() {
	adminTemplates["index"] = loadAdminTemplate("index.gohtml")
	adminTemplates["login"] = loadAdminTemplate("login.gohtml")
}

func loadAdminTemplate(filename string) *template.Template {
	dir := "./templates"
	f1 := path.Join(dir, "layout.gohtml")
	f2 := path.Join(dir, filename)
	return template.Must(template.ParseFiles(f1, f2))
}
