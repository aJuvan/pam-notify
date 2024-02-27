package main

import (
	"fmt"
	"os"

	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/filters"
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

	cont, err := filters.Run(&cfg.Filters, &userData)
	if err != nil {
		config.Logger.Error().Err(err).Msg("Filters failed")
	} else if !cont {
		config.Logger.Debug().Msg("Execution stopped by filters")
		return
	}

	middlewareData := middleware.Run(&cfg.Middlewares, &userData)

	err = notifiers.Run(&cfg.Notifiers, &userData, &middlewareData)
	if err != nil {
		config.Logger.
			Error().
			Err(err).
			Str("service", userData.Service).
			Msg(fmt.Sprintf("Failed to notify user '%s'", userData.Username))
	}
}

func getEnvOrDefault(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
