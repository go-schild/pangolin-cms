package main

import (
	"github.com/go-schild/pangolin-cms/cms"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "cms",
	Short: "Run the cms server.",
	Long:  "Run the cms server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		dataDir, _ := cmd.Flags().GetString("data-dir")
		admin, _ := cmd.Flags().GetBool("admin")

		return cms.Start(cms.Config{
			DataDir:    dataDir,
			AdminPanel: admin,
		})
	},
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initiate CMS environment",
	Long:  "Initiate CMS environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		dataDir, _ := cmd.Flags().GetString("data-dir")
		return cms.Initialize(dataDir)
	},
}

func init() {
	rootCommand.AddCommand(initCommand)

	rootCommand.PersistentFlags().StringP("data-dir", "d", ".", "The data directory.")
	rootCommand.PersistentFlags().Bool("admin", false, "Activate the admin panel")
	_ = rootCommand.MarkFlagRequired("data-dir")
}
