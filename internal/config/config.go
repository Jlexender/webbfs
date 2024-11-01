package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	NetTimeoutSeconds float32 `yaml:"net-timeout-seconds" default:"3"`
	SearchRegexp      string  `yaml:"search-regexp"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
