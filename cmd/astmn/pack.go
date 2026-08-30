package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"astmn/internal/log"
	"astmn/internal/manifest"
	"astmn/internal/packer"
	"astmn/internal/ui"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var newPackage manifest.Manifest

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Create a new astmn package",
	Args:  cobra.ExactArgs(0),

	RunE: func(cmd *cobra.Command, args []string) error {
		log.Infof("launching interactive packer...")

		err := packer.CreateNewPackage(&newPackage)
		if err != nil {
			ui.PError(fmt.Sprintf("failed to create a package: %v", err))
			return err
		}
		log.Infof("new package manifest created!")

		yamlData, err := yaml.Marshal(&newPackage)
		if err != nil {
			ui.PError("failed to marshal manifest to YAML: " + err.Error())
			return err
		}

		manifestFileName := strings.ReplaceAll(strings.TrimSuffix(newPackage.Name, filepath.Ext(newPackage.Name)), " ", "_") + ".yml"
		saveAbsPath, err := filepath.Abs(filepath.Join(c.AssetsRegistryDir, manifestFileName))
		if err != nil {
			ui.PError("unable to get absolute path: " + err.Error())
			return err
		}

		log.Infof("saving manifest file to %s...", saveAbsPath)
		dir := filepath.Dir(saveAbsPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Errorf("failed to create db directory (%v): %v", dir, err)
			return err
		}

		err = os.WriteFile(saveAbsPath, yamlData, 0644)
		if err != nil {
			ui.PError("unable to create/save manifest file: " + err.Error())
			return err
		}

		log.Infof("manifest successfully saved!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(packCmd)
}
