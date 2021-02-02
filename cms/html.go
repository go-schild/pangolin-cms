package cms

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

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
	return template.HTML(data)
}

type HTMLConverter func(config Config, filename string) (template.HTML, error)

var htmlConverters = map[string]HTMLConverter{
	"html": func(config Config, filename string) (template.HTML, error) {
		htmlFile := filepath.Join(config.ContentDir(), "html", filename)
		data, err := ioutil.ReadFile(htmlFile)
		if err != nil {
			return "", err
		}

		return bytesToHTML(data), nil
	},

	"markdown": func(config Config, filename string) (template.HTML, error) {
		mdFile := filepath.Join(config.ContentDir(), "md", filename)
		data, err := ioutil.ReadFile(mdFile)
		if err != nil {
			return "", err
		}

		return bytesToHTML(blackfriday.MarkdownCommon(data)), nil
	},
}
