package main

import (
	"fmt"
	"os"

	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/middleware"
	"github.com/aJuvan/pam-notify/notifiers"
)

func main() {
	pamType := getEnvOrDefault("PAM_TYPE", "unknown")
	if pamType == "close_session" {
		return
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = config.InitLogger(cfg)
	if err != nil {
		panic(err)
	}

	if pamType != "open_session" {
		config.Logger.
			Warn().
			Msg(fmt.Sprintf("Unknown PAM type '%s'", pamType))
	}

	userData := config.UserData{
		Server:   cfg.Server,
		Username: getEnvOrDefault("PAM_USER", "<unknown>"),
		Service:  getEnvOrDefault("PAM_SERVICE", "<unknown>"),
		Rhost:    getEnvOrDefault("PAM_RHOST", "<unknown>"),
	}

	config.Logger.
		Info().
		Str("service", userData.Service).
		Msg(fmt.Sprintf("User '%s' logged in from %s", userData.Username, userData.Rhost))

	middlewareData := middleware.Run(cfg, userData)

	failure := false
	for _, notifier := range cfg.Notifiers {
		f := notifiers.Notifiers[notifier.Type]
		err = f(userData, notifier, middlewareData)
		if err != nil {
			failure = true
			config.Logger.
				Error().
				Str("service", userData.Service).
				Str("notifier", notifier.Type).
				Err(err).
				Msg(fmt.Sprintf("Failed to notify user '%s'", userData.Username))
		}
	}

	if failure {
		os.Exit(1)
	}
}

func getEnvOrDefault(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
