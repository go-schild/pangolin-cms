package cms

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/russross/blackfriday"
)

type PageType string

type Page struct {
	File   string   `yaml:"file"`
	Layout string   `yaml:"layout"`
	CSS    []string `yaml:"css"`
	JS     []string `yaml:"js"`
}

func (p Page) buildHandler(config Config) http.Handler {
	// Load main content
	mdFile := filepath.Join(config.ContentDir, "md", p.File)
	data, _ := ioutil.ReadFile(mdFile) // TODO error handling
	content := bytesToHTML(blackfriday.MarkdownCommon(data))

	// Load layout
	layoutFile := filepath.Join(config.ContentDir, "layouts", p.Layout)
	layout := template.Must(template.ParseFiles(layoutFile))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := layout.Execute(w, map[string]interface{}{
			"Content": content,
			"CSS":     p.CSS,
			"JS":      p.JS,
		})

		if err != nil {
			log.Println(err)
		}
	})
}
