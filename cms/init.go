package cms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func initializeDirs(dataDir string) error {
	fmt.Println("Create basic directory structure")
	defaultDirs := []string{
		"config",
		"content/html",
		"content/layouts",
		"content/md",
		"static/css",
		"static/js",
	}

	for _, d := range defaultDirs {
		joined := path.Join(dataDir, d)
		fmt.Println("Create dir:", joined)
		err := os.MkdirAll(joined, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func initializeFiles(dataDir string) error {
	fmt.Println("Download example files")
	baseURL := `https://raw.githubusercontent.com/go-schild/pangolin-cms/master/example/`
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

		joined := path.Join(dataDir, f)
		fmt.Println("Created:", joined)
		err = ioutil.WriteFile(joined, body, 0600)
		if err != nil {
			return err
		}
	}

	return nil
}

func Initialize(dataDir string) error {
	if err := initializeDirs(dataDir); err != nil {
		return err
	}

	if err := initializeFiles(dataDir); err != nil {
		return err
	}

	return nil
}
