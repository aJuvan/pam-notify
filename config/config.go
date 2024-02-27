package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

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
