package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dev       bool      `yaml:"dev"`
	Debug     bool      `yaml:"debug"`
	Test      string    `yaml:"test"`
	LogConfig LogConfig `yaml:"log,omitempty"`
	WebConfig WebConfig `yaml:"web"`
}

type WebConfig struct {
	Enable bool `yaml:"enable"`
}

type LogConfig struct {
	Level       string
	Filename    string
	MaxSize     int
	MaxAge      int
	MaxBackups  int
	TraceEnable bool   `env:"TRACE_ENABLE" envDefault:"false"`
	TraceUrl    string `env:"TRACE_URL"`
}

// type EnvConfig struct{}

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
