package config

import (
	"os"

	yml "gopkg.in/yaml.v3"
)

type Config struct {
	AssetsRegistryDir string `yaml:"asset_registry_dir"`
	LogsDir           string `yaml:"logs_dir"`
	Preset            string `yaml:"preset"`
}

func (c *Config) GetLogDir() {
	if c.LogsDir == "" {
		return "./logs"
	}
	return c.LogsDir
}

func Load(path string) (*Config, err) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Congig
	if err := yml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
