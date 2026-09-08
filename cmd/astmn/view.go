package main

import (
	"fmt"

	"astmn/internal/log"
	"astmn/internal/manifest"
	"astmn/internal/ui"
	"astmn/internal/viewer"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [manifest.yml]",
	Short: "View the fields of a package manifest file",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		manifestPath := args[0]

		log.Infof("loading manifest (%s)...", manifestPath)
		m, err := manifest.Load(manifestPath)
		if err != nil {
			ui.PError(fmt.Sprintf("failed to load manifest: %v", err))
			return err
		}

		content, err := viewer.RenderManifest(m)
		if err != nil {
			return err
		}

		return viewer.DisplayInPager(content)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
