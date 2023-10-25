package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger(cfg *Config) error {
	file, err := os.OpenFile(
		cfg.Logging.File,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return err
	}

	level, err := getLogLevel(cfg.Logging.Level)
	if err != nil {
		return err
	}

	Logger = zerolog.
		New(file).
		With().Timestamp().
		Logger().
		Level(level)
	return nil
}

func getLogLevel(level string) (zerolog.Level, error) {
	switch level {
	case "debug":
		return zerolog.DebugLevel, nil
	case "info":
		return zerolog.InfoLevel, nil
	case "warn":
		return zerolog.WarnLevel, nil
	case "error":
		return zerolog.ErrorLevel, nil
	case "fatal":
		return zerolog.FatalLevel, nil
	case "panic":
		return zerolog.PanicLevel, nil
	default:
		return zerolog.Disabled, fmt.Errorf("Invalid log level: %s", level)
	}
}
