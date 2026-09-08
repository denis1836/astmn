package viewer

import (
	"fmt"
	"os"
	"path/filepath"

	"astmn/internal/manifest"
	"astmn/internal/ui"
)

func ViewManifest(m *manifest.Manifest) error {
	fmt.Printf(ui.CYellow("Name: ") + m.Name + "\n")
	fmt.Printf(ui.CYellow("Version: ") + m.Version + "\n")
	fmt.Printf(ui.CYellow("Date: ") + m.Date + "\n")

	fmt.Printf(ui.CYellow("Author: ") + m.Author + "\n")
	fmt.Printf(ui.CYellow("Description: ") + m.Description + "\n")
	fmt.Printf(ui.CYellow("Contributors: ") + "\n")
	for _, c := range m.Contributors {
		fmt.Printf(" %s\n", c)
	}
	fmt.Println()

	fmt.Printf(ui.CYellow("File name: ") + m.FileName + "\n")
	fmt.Printf(ui.CYellow("Download URL: ") + m.DownloadURL + "\n")
	fmt.Printf(ui.CYellow("SHA256: ") + m.SHA256 + "\n")

	ex, err := os.Executable()
	if err != nil {
		return err
	}
	runDir := filepath.Dir(ex)

	insPathRelDir, err := filepath.Rel(runDir, ex)
	if err != nil {
		return err
	}
	fmt.Printf(ui.CYellow("InstallPath: ") + insPathRelDir + "/" + m.InstallPath + "\n")

	fmt.Printf(ui.CYellow("Contents: ") + "\n")
	for _, c := range m.Contents {
		fmt.Printf(" %s\n", c)
	}
	fmt.Println()
	fmt.Printf(ui.CYellow("Dependencies: ") + "\n")
	for _, d := range m.DependsOn {
		fmt.Printf(" %s\n", d)
	}
	fmt.Println()
	fmt.Printf(ui.CYellow("Changelog: ") + "\n")
	for _, ch := range m.Changelog {
		fmt.Printf(" %s\n", ch)
	}
	fmt.Println()

	return nil
}
