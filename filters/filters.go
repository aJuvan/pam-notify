package filters

import (
	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/filters/ipfilter"
)

func Run(conf *config.ConfigFilters, userData *config.UserData) (bool, error) {
	filtersLogger := config.Logger.With().Str("module", "filters").Logger()

	continueExecution := true

	if conf.IPFilter.Enabled {
		filtersLogger.Debug().Msg("Running IP filters")
		cont, err := ipfilter.Run(&conf.IPFilter, userData)
		if err != nil {
			filtersLogger.Error().Err(err).Msg("IP filters failed")
			return true, err
		}
		continueExecution = continueExecution && cont
	}

	return continueExecution, nil
}
