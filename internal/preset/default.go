package preset

import (
	"astmn/internal/manifest"
)

type DefaultPreset struct{}

func (p *DefaultPreset) Name() string { return "ue5" }

func (p *DefaultPreset) Validate(m *manifest.Manifest) []string {
	var errs []string

	//TODO
}
