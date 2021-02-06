package cms

import (
	"html/template"
)

var adminTemplates = map[string]*template.Template{}

func initAdminTemplates() {
	adminTemplates["index"] = loadAdminTemplate("admin/index.gohtml", "admin/layout.gohtml")
	adminTemplates["login"] = loadAdminTemplate("admin/login.gohtml", "admin/layout.gohtml")
}

func loadAdminTemplate(filename string, layout string) *template.Template {
	content1, _ := templateBox.FindString(filename)
	content2, _ := templateBox.FindString(layout)

	t := template.New("")
	t = template.Must(t.Parse(content1))
	t = template.Must(t.Parse(content2))

	return t
}

func handleAdminIndex(c *Context) error {
	t := adminTemplates["index"]
	return t.Execute(c.Writer, nil)
}

func handleAdminLogin(c *Context) error {
	t := adminTemplates["login"]
	return t.Execute(c.Writer, nil)
}
