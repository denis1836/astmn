package preset

import (
	"fmt"
	"path/filepath"

	"astmn/internal/manifest"
)

type FreeCADPreset struct{}

func (p *FreeCADPreset) Name() string { return "freecad" }

func (p *FreeCADPreset) Validate(m *manifest.Manifest) []string {
	var errs []string = nil

	//TODO: add freecad-specific fields

	return errs
}

func (p *FreeCADPreset) ResolveTargetDir(baseDir, installPath string) (string, error) {
	if installPath == "" {
		return "", fmt.Errorf("manifest install path is missing")
	}
	return filepath.Join(baseDir, installPath), nil
}

func (p *FreeCADPreset) InspectArchive(archivePath string) error {
	//TODO

	return nil
}

func (p *FreeCADPreset) PostInstall(targetDir string, m *manifest.Manifest) error {
	//TODO: give execute permissions to python scripts

	return nil
}
