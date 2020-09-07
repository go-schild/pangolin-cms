package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-schild/pangolin-cms/cms"
)

var (
	configDir  string
	contentDir string
	staticDir  string
)

func init() {
	flag.StringVar(&configDir, "config", "./config", "The configuration directory")
	flag.StringVar(&contentDir, "content", "./content", "The content directory")
	flag.StringVar(&staticDir, "static", "./static", "The directory with the static files")
	flag.Parse()
}

func main() {
	err := cms.Start(cms.Config{
		ConfigDir:  configDir,
		ContentDir: contentDir,
		StaticDir:  staticDir,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
