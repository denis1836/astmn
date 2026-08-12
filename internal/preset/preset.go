package preset

import (
	"astmn/internal/manifest"	
)

type PresetHandler interface {
	Name() string

	Validate(m *manifest.Manifest) []string

	ResolveTargetDir(baseDir, manifestInstallPath string) string

	InspectArchive(archivePath string) error

	PostInstall(targetDir string, m *manifest.Manifest) error
}

func Get(name string, defaultPreset string) PresetHandler {
	if name == "" {
		name = defaultPreset
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
