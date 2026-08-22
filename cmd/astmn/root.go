package main

import (
	"fmt"
	"os"

	"astmn/internal/opts"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "astmn",
	Short: "astmn - Asset Manager for game dev, CAD and large binary files",
	Long:  `Cross-platform CLI tool for managing large binary assets using lightweight YAML manifests and SQLite.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&opts.App.Verbose, "verbose", "v", false, "Enable verbose output logging")
}
