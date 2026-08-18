package preset

import (
	"astmn/internal/manifest"
	"fmt"
	"path/filepath"
)

type DefaultPreset struct{}

func (p *DefaultPreset) Name() string { return "default" }

func (p *DefaultPreset) Validate(m *manifest.Manifest) []string {
	var errs []string

	if m.Name == "" {
		errs = append(errs, "name is missing")
	}
	if m.Version == "" {
		errs = append(errs, "verstion is missing")
	}
	if m.DownloadURL == "" {
		errs = append(errs, "download_url is missing")
	}

	return errs
}

func (p *DefaultPreset) ResolveTargetDir(baseDir, installPath string) (string, error) {
	if installPath == "" {
		return "", fmt.Errorf("manifest install path is missing")
	}
	return filepath.Join(baseDir, installPath), nil
}

func (p *DefaultPreset) InspectArchive(archivePath string) error {
	//TODO
	return nil
}

func (p *DefaultPreset) PostInstall(targetDir string, m *manifest.Manifest) error {
	return nil
}
