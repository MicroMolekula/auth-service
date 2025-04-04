package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server        `yaml:"server"`
	Database      `yaml:"database"`
	JWT           `yaml:"jwt"`
	SessionSecret string `yaml:"session_secret"`
	OauthYandex   `yaml:"oauth_yandex"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Timezone string `yaml:"timezone"`
}

type OauthYandex struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURl  string `yaml:"redirect_url"`
	State        string `yaml:"state"`
}

type JWT struct {
	Secret     string `yaml:"secret"`
	TTL        int    `yaml:"ttl"`
	RefreshTTL int    `yaml:"refresh_ttl"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	yamlPath := os.Getenv("CONFIG_PATH")
	if yamlPath == "" {
		return nil, errors.New("config file environment variable not set")
	}
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error read config file: %s", err))
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parse config file: %s", err))
	}
	return cfg, nil
}
