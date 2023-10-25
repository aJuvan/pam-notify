package middleware

import (
	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/middleware/geoip"
)

type MiddlewareData struct {
	GeoIP *geoip.MiddlewareGeoIPData
}

func Run(conf *config.Config, userData config.UserData) MiddlewareData {
	middlewareLogger := config.Logger.With().Str("module", "middleware").Logger()
	middlewareData := MiddlewareData{}

	if conf.Middlewares == nil {
		return middlewareData
	}

	ch := make(chan error)
	routines := 0

	if conf.Middlewares.GeoIP != nil && *conf.Middlewares.GeoIP {
		routines++
		go func() {
			geoIPData, err := geoip.Run(&userData)
			if err != nil {
				middlewareLogger.
					Error().
					Err(err).
					Msg("Error while running geoip middleware")
				ch <- err
				return
			}
			middlewareData.GeoIP = geoIPData
			ch <- nil
		}()
	}

	failures := 0
	for i := 0; i < routines; i++ {
		err := <-ch
		if err != nil {
			failures += 1
		}
	}

	middlewareLogger.
		Info().
		Msgf("Ran %d out of %d middleware rutined", routines-failures, routines)

	return middlewareData
}
