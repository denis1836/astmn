package manifest

import (
	"os"

	yml "gopkg.in/yaml.v3"
)

type Manifest struct {
	Name string `yaml:"name"`
	Version string `yaml:"version"`
	Date string `yaml:"date"`
	
	Author string `yaml:"author"`
	Description string `yaml:"description"`
	Contributors [] string `yaml:"contributors"`

	FileName string `yaml:"file_name"`
	DownloadURL string `yaml:"download_url"`
	SHA256 string `yaml:"sha256"`
	ArchiveType string `yaml:"archive_type"`
	InstallPath string `yaml:"install_path"`

	Contents []string `yaml:"contents"`
	DependsOn []string `yaml:"depends_on"`
	Changelog []string `yaml:"changelog"`
}

func Load(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var m Manifest
	if err := yml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	 
	return &m, nil
}
