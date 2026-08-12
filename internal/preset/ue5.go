package preset

import (
	"fmt"
	"path/filepath"

	"astmn/internal/ui"	
	"astmn/internal/manifest"
)

type UE5Preset struct {}

func (p *UE5Preset) Name() string { return "ue5" }

func (p *UE5Preset) Validate(m* manifest.Manifest) []string {
	var errs []string
	
	//TODO
}
