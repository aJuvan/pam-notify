package notifiers

import (
	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/middleware"
	"github.com/aJuvan/pam-notify/notifiers/discord"
)

type Notifier func(config.UserData, config.ConfigNotifier, middleware.MiddlewareData) error

var Notifiers = map[string]Notifier{
	"discord": discord.Notify,
}
