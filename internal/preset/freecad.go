package preset

import (
	"astmn/internal/manifest"
)

type FreeCADPreset struct{}

func (p *FreeCADPreset) Name() string { return "ue5" }

func (p *FreeCADPreset) Validate(m *manifest.Manifest) []string {
	var errs []string

	//TODO
}
