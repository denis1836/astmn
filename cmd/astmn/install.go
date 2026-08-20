package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"astmn/internal/downloader"
	"astmn/internal/log"
	"astmn/internal/manifest"
	"astmn/internal/ui"
)

var (
	forceFlag    bool
	downloadOnly bool
)

var installCmd = &cobra.Command{
	Use:   "install [manifest.yml]",
	Short: "Install an asset package from a manifest file",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		manifestPath := args[0]

		log.Infof("loading manifest (%s)...", manifestPath)
		m, err := manifest.Load(manifestPath)
		if err != nil {
			ui.PError(fmt.Sprintf("failed to load manifest: %v", err))
			return err
		}

		ui.PInfo(fmt.Sprintf("installing package: %s(%v)", m.Name, m.Version))
		log.Infof("starting download (%s)...", m.Name)

		//TODO: validating/clearing dir from config for safety
		destArchive := fmt.Sprintf(c.TempDownloadDir + "/" + m.FileName)
		if err := downloader.DownloadFile(m.DownloadURL, destArchive); err != nil {
			ui.PError(fmt.Sprintf("failed to download: %w", err))
			return err
		}

		if downloadOnly {
			ui.POk("download completed (--download-only active), skipping extraction")
			return nil
		}

		//TODO: extractor module
		//TODO: verification (extensions, zip searching)
		//TODO: db module actions
		//TODO: post install preset actions

		ui.POk(fmt.Sprintf("successfully installed %s", m.Name))
		return nil
	},
}

func init() {
	installCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force re-installation even if package exists")
	installCmd.Flags().BoolVar(&downloadOnly, "download-only", false, "Download to archive without extracting")

	rootCmd.AddCommand(installCmd)
}
