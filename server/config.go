package server

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

const (
	cfgFilePath = "./configuration.yaml"
)

type Config struct {
	Address      string        `yaml:"address,omitempty"`
	ReadTimeout  time.Duration `yaml:"read-timeout,omitempty"`
	WriteTimeout time.Duration `yaml:"write-timeout,omitempty"`
}

func NewConfig() (*Config, error) {
	cfg, err := loadConfigFromFile()
	if err != nil {
		return nil, err
	}
	cfg.setEmptyCfgByDefaults()

	return cfg, nil
}

func loadConfigFromFile() (*Config, error) {
	content, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		return nil, err
	}

	cfg := Config{}

	if err = yaml.Unmarshal(content, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) setEmptyCfgByDefaults() {
	if c.Address == "" {
		c.Address = "localhost:8080"
	}

	if c.ReadTimeout == 0*time.Second {
		c.ReadTimeout = time.Second * 10
	}

	if c.WriteTimeout == 0*time.Second {
		c.WriteTimeout = time.Second * 20
	}
}
