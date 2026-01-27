package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server" json:"server"`
	Download DownloadConfig `yaml:"download" json:"download"`
}

type ServerConfig struct {
	Port int `yaml:"port" json:"port"`
}

type DownloadConfig struct {
	Path string     `yaml:"path" json:"path"`
	Bing BingConfig `yaml:"bing" json:"bing"`
}

type BingConfig struct {
	Enabled  bool   `yaml:"enabled" json:"enabled"`
	Schedule string `yaml:"schedule" json:"schedule"`
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

// SaveConfig 保存配置到指定路径
func SaveConfig(path string, config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
