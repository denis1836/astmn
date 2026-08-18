package preset

import (
	"fmt"
	"path/filepath"

	"astmn/internal/manifest"
)

type UE5Preset struct{}

func (p *UE5Preset) Name() string { return "ue5" }

func (p *UE5Preset) Validate(m *manifest.Manifest) []string {
	var errs []string = nil

	if m.Name == "" {
		errs = append(errs, "name is missing")
	}
	if m.Version == "" {
		errs = append(errs, "verstion is missing")
	}
	if m.DownloadURL == "" {
		errs = append(errs, "download_url is missing")
	}
	//TODO

	return nil
}

func (p *UE5Preset) ResolveTargetDir(baseDir, installPath string) (string, error) {
	if installPath == "" {
		return "", fmt.Errorf("manifest install path is missing")
	}
	return filepath.Join(baseDir, installPath), nil
}

func (p *UE5Preset) InspectArchive(archivePath string) error {
	//TODO

	return nil
}

func (p *UE5Preset) PostInstall(targetDir string, m *manifest.Manifest) error {
	//TODO: give execute permissions to python scripts

	return nil
}
