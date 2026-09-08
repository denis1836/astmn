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

	fmt.Fprintf(&buf, ui.CYellow("Name: ")+m.Name+"\n")
	fmt.Fprintf(&buf, ui.CYellow("Version: ")+m.Version+"\n")
	fmt.Fprintf(&buf, ui.CYellow("Date: ")+m.Date+"\n")

	fmt.Fprintf(&buf, ui.CYellow("Author: ")+m.Author+"\n")
	fmt.Fprintf(&buf, ui.CYellow("Description: ")+m.Description+"\n")
	fmt.Fprintf(&buf, ui.CYellow("Contributors: ")+"\n")
	for _, c := range m.Contributors {
		fmt.Fprintf(&buf, " %s\n", c)
	}
	fmt.Fprintln(&buf)

	fmt.Fprintf(&buf, ui.CYellow("File name: ")+m.FileName+"\n")
	fmt.Fprintf(&buf, ui.CYellow("Download URL: ")+m.DownloadURL+"\n")
	fmt.Fprintf(&buf, ui.CYellow("SHA256: ")+m.SHA256+"\n")

	ex, err := os.Executable()
	if err != nil {
		return buf.String(), err
	}
	runDir := filepath.Dir(ex)

	insPathRelDir, err := filepath.Rel(runDir, ex)
	if err != nil {
		return buf.String(), err
	}
	fmt.Fprintf(&buf, ui.CYellow("InstallPath: ")+insPathRelDir+"/"+m.InstallPath+"\n")

	fmt.Fprintf(&buf, ui.CYellow("Contents: ")+"\n")
	for _, c := range m.Contents {
		fmt.Fprintf(&buf, " %s\n", c)
	}
	fmt.Fprintln(&buf)

	fmt.Fprintf(&buf, ui.CYellow("Dependencies: ")+"\n")
	for _, d := range m.DependsOn {
		fmt.Fprintf(&buf, " %s\n", d)
	}
	fmt.Fprintln(&buf)

	fmt.Fprintf(&buf, ui.CYellow("Changelog: ")+"\n")
	for _, ch := range m.Changelog {
		fmt.Fprintf(&buf, " %s\n", ch)
	}
	fmt.Fprintln(&buf)

	return buf.String(), nil
}
