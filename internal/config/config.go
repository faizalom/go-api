package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// App holds the loaded configuration
var App Config

// Config defines the structure of the configuration file
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

// Load reads the configuration file from the given path and unmarshals it.
func Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &App); err != nil {
		return err
	}

	return nil
}
