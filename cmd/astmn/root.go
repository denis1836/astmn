package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verboseFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "astmn",
	Short: "astmn - Asset Manager for gamedev, CAD and large binary files",
	Long:  `Cross-platform CLI tool for managing large binary assets using lightweight YAML manifests and SQLite.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//TODO: configure verbose flag functionality for other modules
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Enable verbose output logging")
}
