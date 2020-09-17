package cms

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type PageType string

type Page struct {
	Index    string        `yaml:"-"`
	URL      string        `yaml:"url"`
	Title    string        `yaml:"title"`
	Layout   string        `yaml:"layout"`
	Contents []PageContent `yaml:"content"`
	CSS      []string      `yaml:"css"`
	JS       []string      `yaml:"js"`
}

type PageContent struct {
	Type  string `yaml:"type"`
	File  string `yaml:"file"`
	ID    string `yaml:"id"`
	Class string `yaml:"class"`
}

func (p Page) buildHandler(config Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var contents []templateDataContentStruct

		// Load layout
		layoutFile := filepath.Join(config.ContentDir, "layouts", p.Layout)
		layout := template.Must(template.ParseFiles(layoutFile))

		// Load contents
		for _, pageContent := range p.Contents {
			converter, ok := htmlConverters[pageContent.Type]
			if !ok {
				// TODO error or something.
				continue
			}

			html, _ := converter(config, pageContent.File)  // TODO error handling
			contents = append(contents, templateDataContentStruct{
				HTML: html,
				ID: pageContent.ID,
				Class: pageContent.Class,
			})
		}

		err := layout.Execute(w, templateDataStruct{
			Title:    p.Title,
			Contents: contents,
			CSS:      p.CSS,
			JS:       p.JS,
		})

		if err != nil {
			log.Println(err)
		}
	})
}

func loadPageList(url string) ([]Page, error) {
	var list []Page             // The result
	var mapList map[string]Page // The structure of the file
	err := unmarshal(url, &mapList)
	if err != nil {
		return nil, err
	}

	// Key of the map is the index, write it into the field "Index"
	for index, page := range mapList {
		page.Index = index
		list = append(list, page)
	}

	return list, nil
}
