package main

import (
	"github.com/go-schild/pangolin-cms/cms"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "cms",
	Short: "Run the cms server.",
	Long: "Run the cms server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cms.Start(cms.Config{
			ConfigDir:  configDir,
			ContentDir: contentDir,
			StaticDir:  staticDir,
		})
	},
}

var initCommand = &cobra.Command{
	Use: "init",
	Short: "Initiate CMS environment",
	Long: "Initiate CMS environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cms.Initialize()
	},
}

func init() {
	rootCommand.AddCommand(initCommand)
}
