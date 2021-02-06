package cms

import (
	"net/http"
	"path/filepath"
)

func Start(config Config) error {
	config.fillDefaultValues()
	server := NewServer()

	// Load pages
	pagesConfigFile := filepath.Join(config.ConfigDir(), "pages.yml")
	pages, err := loadPageList(pagesConfigFile)
	if err != nil {
		return err
	}

	// Build handlers
	for _, urlInfo := range pages {
		handler := urlInfo.buildHandler(config)
		server.handle(urlInfo.URL, handler)
	}

	// Admin
	if config.AdminPanel {
		initAdminTemplates()
		server.handle("/admin/", handleAdminIndex)
		server.handle("/admin/login/", handleAdminLogin)
	}

	// Static
	staticHandler := http.FileServer(http.Dir(config.StaticDir()))
	server.mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	return server.Run(":8080")
}
