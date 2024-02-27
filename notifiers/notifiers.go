package notifiers

import (
	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/middleware"
	"github.com/aJuvan/pam-notify/notifiers/discord"
)

func Run(conf *config.ConfigNotifiers, userData *config.UserData, middlewareData *middleware.MiddlewareData) error {
	notifiersLogger := config.Logger.With().Str("module", "notifiers").Logger()

	for i, d := range conf.Discord {
		notifiersLogger.Debug().Int("index", i).Msg("Running Discord notifier")
		err := discord.Run(&d, userData, middlewareData)
		if err != nil {
			notifiersLogger.Error().Err(err).Int("index", i).Msg("Discord notifier failed")
			return err
		}
	}

	return nil
}
