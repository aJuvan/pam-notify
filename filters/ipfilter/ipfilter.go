package ipfilter

import (
	"net"

	"github.com/aJuvan/pam-notify/config"
)

func Run(conf *config.ConfigFiltersIPFilter, userData *config.UserData) (bool, error) {
	filtersLogger := config.Logger.With().Str("module", "filters").Logger()

	for _, cidr := range conf.Cidrs {
		filtersLogger.Debug().Str("cidr", cidr).Msg("Checking CIDR")
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			filtersLogger.Error().Err(err).Str("cidr", cidr).Msg("Invalid CIDR")
			return true, err
		}
		if ipnet.Contains(net.ParseIP(userData.Rhost)) {
			filtersLogger.Info().Str("cidr", cidr).Msg("IP allowed")
			return false, nil
		}
	}

	return true, nil
}
