package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"astmn/internal/db"
	"astmn/internal/downloader"
	"astmn/internal/extractor"
	"astmn/internal/log"
	"astmn/internal/manifest"
	"astmn/internal/preset"
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

		//TODO: validator package

		destArchive := fmt.Sprintf("%s/%s", c.TempDownloadDir, m.FileName)
		log.Infof("starting download to %s...", destArchive)
		if err := downloader.DownloadFile(m.DownloadURL, destArchive); err != nil {
			ui.PError(fmt.Sprintf("failed to download: %w", err))
			return err
		}

		pkgID, err := db.InsertPackage(m.Name, m.Version, c.Preset)
		if err != nil {
			ui.PError(fmt.Sprintf("failed to register package in database: %v", err))
			return err
		}

		if downloadOnly {
			ui.POk("download completed (--download-only active), skipping extraction")
			return nil
		}

		ui.PInfo("extracting files...")
		extractedFiles, err := extractor.ExtractArchive(destArchive, m.InstallPath)
		if err != nil {
			ui.PError(fmt.Sprintf("failed to extract archive: %v", err))
			return err
		}

		log.Infof("logging %d extracted files to db...", len(extractedFiles))
		for _, file := range extractedFiles {
			if err := db.InsertFile(pkgID, file.Name, file.RelPath, file.Hash, file.Size); err != nil {
				log.Errorf("failed to log file %s to db: %v", file.RelPath, err)
			}
		}

		pHandler := preset.Get(c.Preset)
		if pHandler != nil {
			log.Infof("executing post-install actions for preset: %s", pHandler.Name())
			if err := pHandler.PostInstall(m.InstallPath, m); err != nil {
				ui.PError(fmt.Sprintf("post-install hook failed: %v", err))
				return err
			}
		}

		ui.POk(fmt.Sprintf("successfully installed %s", m.Name))
		return nil
	},
}

func init() {
	installCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force re-installation even if package exists")
	installCmd.Flags().BoolVar(&downloadOnly, "download-only", false, "Download to archive without extracting")

	rootCmd.AddCommand(installCmd)
}
