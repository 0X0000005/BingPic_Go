package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Download DownloadConfig `yaml:"download"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DownloadConfig struct {
	Path      string          `yaml:"path"`
	Bing      BingConfig      `yaml:"bing"`
	Spotlight SpotlightConfig `yaml:"spotlight"`
}

type BingConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Schedule string `yaml:"schedule"`
}

type SpotlightConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Schedule string `yaml:"schedule"`
}

// LoadConfig 从指定路径加载 yaml 配置
func LoadConfig(path string) (*Config, error) {
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
