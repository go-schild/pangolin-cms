package cms

import "html/template"

func bytesToHTML(data []byte) template.HTML {
	return template.HTML(string(data))
}
