package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Timeout      float32 `yaml:"timeout"`
	SearchRegexp string  `yaml:"search-regexp"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
