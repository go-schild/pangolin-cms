package cms

import "html/template"

type templateDataStruct struct {
	Title    string
	Contents []template.HTML
	CSS      []string
	JS       []string
}

func bytesToHTML(data []byte) template.HTML {
	return template.HTML(string(data))
}
