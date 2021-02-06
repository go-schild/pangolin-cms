package cms

func handleAdminIndex(c *Context) error {
	t := adminTemplates["index"]
	return t.Execute(c.Writer, nil)
}

func handleAdminLogin(c *Context) error {
	t := adminTemplates["login"]
	return t.Execute(c.Writer, nil)
}
