package config

import (
	"os"

	yml "gopkg.in/yaml.v3"
)

type Config struct {
	AssetsRegistryDir string `yaml:"asset_registry_dir"`
	LogsDir           string `yaml:"logs_dir"`
	DBPath            string `yaml:"db_dir"`
	Preset            string `yaml:"preset"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
