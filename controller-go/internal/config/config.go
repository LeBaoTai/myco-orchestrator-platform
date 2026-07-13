package config

import (
	"os"

	"go.yaml.in/yaml/v4"
)

type SchemaConfig struct {
	Repository  string   `yaml:"repository"`
	SearchPaths []string `yaml:"searchPaths"`
	Modules     []string `yaml:"modules"`
}

type Config struct {
	Schema SchemaConfig `yaml:"schema"`
}

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
