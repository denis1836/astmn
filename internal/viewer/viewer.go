package viewer

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"astmn/internal/manifest"
	"astmn/internal/ui"
)

func RenderManifest(m *manifest.Manifest) (string, error) {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "%s%s\n", ui.CYellow("Name: "), m.Name)
	fmt.Fprintf(&buf, "%s%s\n", ui.CYellow("Version: "), m.Version)
	fmt.Fprintf(&buf, "%s%s\n", ui.CYellow("Date: "), m.Date)

	fmt.Fprintf(&buf, "%s%s\n", ui.CYellow("Author: "), m.Author)
	fmt.Fprintf(&buf, "%s%s\n", ui.CYellow("Description: "), m.Description)

	fmt.Fprintln(&buf, ui.CYellow("Contributors: "))
	for _, c := range m.Contributors {
		fmt.Fprintf(&buf, " %s\n", c)
	}
	fmt.Fprintln(&buf)

	fmt.Fprintf(&buf, "%s%s\n", ui.CYellow("File name: "), m.FileName)
	fmt.Fprintf(&buf, "%s%s\n", ui.CYellow("Download URL: "), m.DownloadURL)
	fmt.Fprintf(&buf, "%s%s\n", ui.CYellow("SHA256: "), m.SHA256)

	ex, err := os.Executable()
	if err != nil {
		return buf.String(), err
	}
	runDir := filepath.Dir(ex)

	insPathRelDir, err := filepath.Rel(runDir, ex)
	if err != nil {
		return buf.String(), err
	}
	fmt.Fprintf(&buf, "%s%s/%s\n", ui.CYellow("InstallPath: "), insPathRelDir, m.InstallPath)

	fmt.Fprintln(&buf, ui.CYellow("Contents: "))
	for _, c := range m.Contents {
		fmt.Fprintf(&buf, " %s\n", c)
	}
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, ui.CYellow("Dependencies: "))
	for _, d := range m.DependsOn {
		fmt.Fprintf(&buf, " %s\n", d)
	}
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, ui.CYellow("Changelog: "))
	for _, ch := range m.Changelog {
		fmt.Fprintf(&buf, " %s\n", ch)
	}
	fmt.Fprintln(&buf)

	return buf.String(), nil
}
