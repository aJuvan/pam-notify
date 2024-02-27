package middleware

import (
	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/middleware/geoip"
)

type MiddlewareData struct {
	GeoIP *geoip.MiddlewareGeoIPData
}

func Run(conf *config.ConfigMiddlewares, userData *config.UserData) MiddlewareData {
	middlewareLogger := config.Logger.With().Str("module", "middleware").Logger()
	middlewareData := MiddlewareData{}

	if conf.GeoIP.Enabled {
		middlewareLogger.Debug().Msg("Running GeoIP middleware")
		tmp, err := geoip.Run(&conf.GeoIP, userData)
		if err != nil {
			middlewareLogger.Error().Err(err).Msg("GeoIP middleware failed")
		} else {
			middlewareData.GeoIP = tmp
		}
	}

	return middlewareData
}
