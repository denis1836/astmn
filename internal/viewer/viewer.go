package viewer

import (
	"fmt"
	"os"
	"path/filepath"

	"astmn/internal/manifest"
)

func ViewManifest(m *manifest.Manifest) error {
	fmt.Printf("Name: %s\n", m.Name)
	fmt.Printf("Version: %s\n", m.Version)
	fmt.Printf("Date: %s\n", m.Date)

	fmt.Printf("Author: %s\n", m.Author)
	fmt.Printf("Description: %s\n", m.Description)
	fmt.Printf("Contributors: \n")
	for _, c := range m.Contributors {
		fmt.Printf(" %s\n", c)
	}
	fmt.Println()

	fmt.Printf("File name: %s\n", m.FileName)
	fmt.Printf("Download URL: %s\n", m.DownloadURL)
	fmt.Printf("SHA256: %s\n", m.SHA256)

	ex, err := os.Executable()
	if err != nil {
		return err
	}
	runDir := filepath.Dir(ex)

	insPathRelDir, err := filepath.Rel(runDir, ex)
	if err != nil {
		return err
	}
	fmt.Printf("InstallPath: %s\n", insPathRelDir+"/"+m.InstallPath)

	fmt.Printf("Contents: \n")
	for _, c := range m.Contents {
		fmt.Printf(" %s\n", c)
	}
	fmt.Println()
	fmt.Printf("Dependencies: \n")
	for _, d := range m.DependsOn {
		fmt.Printf(" %s\n", d)
	}
	fmt.Println()
	fmt.Printf("Changelog: \n")
	for _, ch := range m.Changelog {
		fmt.Printf(" %s\n", ch)
	}
	fmt.Println()

	return nil
}
