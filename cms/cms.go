package cms

import (
	"net/http"
	"path/filepath"
)

func Start(config Config) error {
	config.fillDefaultValues()

	// Load urls
	urlsConfigFile := filepath.Join(config.ConfigDir, "urls.yml")
	urlsSet, err := loadWebURLSet(urlsConfigFile)
	if err != nil {
		return err
	}

	// Build handlers
	for _, urlInfo := range urlsSet {
		handler := urlInfo.Page.buildHandler(config)
		http.Handle(urlInfo.URL, handler)
	}

	// Build static handler
	staticHandler := http.FileServer(http.Dir(config.StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	return http.ListenAndServe(config.WebAddr, nil)
}
