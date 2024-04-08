package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dev       bool      `yaml:"dev"`
	Port      string    `yaml:"port"`
	WebEnable bool      `yaml:"web_enable"`
	LogConfig LogConfig `yaml:"log,omitempty"`
}

type LogConfig struct {
	Level       string
	Filename    string `yaml:"filename"`
	MaxSize     int
	MaxAge      int
	MaxBackups  int
	TraceEnable bool
	TraceUrl    string
}

func NewConfig(path string) (*Config, error) {
	var cfg = Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
