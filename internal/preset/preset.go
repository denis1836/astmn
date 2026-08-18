package preset

import (
	"astmn/internal/manifest"
)

type PresetHandler interface {
	Name() string

	Validate(m *manifest.Manifest) []string

	ResolveTargetDir(baseDir, manifestInstallPath string) (string, error)

	InspectArchive(archivePath string) error

	PostInstall(targetDir string, m *manifest.Manifest) error
}

func Get(name string) PresetHandler {
	if name == "" {
		name = "default"
	}

	switch name {
	case "ue5":
		return &UE5Preset{}
	case "freecad":
		return &FreeCADPreset{}
	default:
		return &DefaultPreset{}
	}
}
