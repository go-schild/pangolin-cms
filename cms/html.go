package cms

import "html/template"

type templateDataStruct struct {
	Title    string
	Contents []templateDataContentStruct
	CSS      []string
	JS       []string
}

type templateDataContentStruct struct {
	HTML  template.HTML
	ID    string
	Class string
}

func bytesToHTML(data []byte) template.HTML {
	return template.HTML(string(data))
}
