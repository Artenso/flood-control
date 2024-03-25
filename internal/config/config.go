package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TimeInterval int64  `yaml:"time_interval"`
	CallsLimit   int64  `yaml:"calls_limit"`
	RedisURL     string `yaml:"redis_url"`
}

func Read(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
