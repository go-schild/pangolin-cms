package cms

import (
	"net/http"
	"path/filepath"
)

func Start(config Config) error {
	config.fillDefaultValues()

	// Load pages
	pagesConfigFile := filepath.Join(config.ConfigDir, "pages.yml")
	pages, err := loadPageList(pagesConfigFile)
	if err != nil {
		return err
	}

	// Build handlers
	for _, urlInfo := range pages {
		handler := urlInfo.buildHandler(config)
		http.Handle(urlInfo.URL, handler)
	}

	// Build static handler
	staticHandler := http.FileServer(http.Dir(config.StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	return http.ListenAndServe(config.WebAddr, nil)
}
