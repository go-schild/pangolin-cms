package cms

import (
	"io/ioutil"
	"net/http"
	"os"
)

func initializeDirs() error {
	defaultDirs := []string{
		"config",
		"content/html",
		"content/layouts",
		"content/md",
		"static/css",
		"static/js",
	}

	for _, d := range defaultDirs {
		err := os.MkdirAll(d, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func initializeFiles() error {
	// TODO isn't there a more elegant way?
	baseURL := `https://raw.githubusercontent.com/go-schild/pangolin-cms/master/cmd/cms/`
	copyFiles := []string{
		"config/pages.yml",
		"content/html/index.html",
		"content/layouts/layout.gohtml",
		"content/md/config.md",
		"content/md/contents.md",
		"content/md/index.md",
		"content/md/starting.md",
		"static/css/index.less",
		"static/css/index.css",
		"static/css/main.less",
		"static/css/main.css",
	}

	for _, f := range copyFiles {
		resp, err := http.Get(baseURL + f)
		if err != nil {
			return err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(f, body, 0600)
		if err != nil {
			return err
		}
	}

	return nil
}

func Initialize() error {
	if err := initializeDirs(); err != nil {
		return err
	}

	if err := initializeFiles(); err != nil {
		return err
	}

	return nil
}
