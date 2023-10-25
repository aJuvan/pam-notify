package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type UserData struct {
	Server   string
	Username string
	Service  string
	Rhost    string
}

type Config struct {
	Server      string            `json:"hostname"`
	Logging     ConfigLogging     `json:"logging"`
	Notifiers   []ConfigNotifier  `json:"notifiers"`
	Middlewares *MiddlewareConfig `json:"middlewares"`
}

type ConfigLogging struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

type ConfigNotifier struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

func LoadConfig() (*Config, error) {
	configFile := flag.String("config", "/etc/pam-notify.yml", "Path to config file")
	flag.Parse()

	file, err := os.Open(*configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
