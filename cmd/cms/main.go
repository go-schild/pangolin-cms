package main

import (
	"flag"
	"fmt"
	"os"
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
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
