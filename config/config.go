package config

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	// ErrorConfigFileNotFound error
	ErrorConfigFileNotFound error = errors.New("configuration file not found")
	// ErrorConfigMissingRoom error
	ErrorConfigMissingRoom error = errors.New("configuration is missing room")
	// ErrorConfigInputDeviceMissing error
	ErrorConfigInputDeviceMissing error = errors.New("configuration is missing input device")
)

// Config struct
type Config struct {
	Room        string `json:"room" yaml:"room"`
	InputDevice string `json:"input_device" yaml:"input_device"`
}

// Load returns a pointer to Config
func Load(filename string) (*Config, error) {
	var cfg Config

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	if cfg.Room == "" {
		return nil, ErrorConfigMissingRoom
	}

	if cfg.InputDevice == "" {
		return nil, ErrorConfigInputDeviceMissing
	}
	
	return &cfg, nil
}

