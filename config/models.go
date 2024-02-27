package config

type UserData struct {
	Server   string
	Username string
	Service  string
	Rhost    string
}

type Config struct {
	Server      string            `json:"server"`
	Logging     ConfigLogging     `json:"logging"`
	Notifiers   ConfigNotifiers   `json:"notifiers"`
	Middlewares ConfigMiddlewares `json:"middlewares"`
	Filters     ConfigFilters     `json:"filters"`
}

type ConfigLogging struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

type ConfigNotifiers struct {
	Discord  []ConfigNotifierDiscord  `json:"discord"`
	Telegram []ConfigNotifierTelegram `json:"telegram"`
}

type ConfigNotifierDiscord struct {
	Webhook string `json:"webhook"`
}

type ConfigNotifierTelegram struct {
	ChatId string `json:"chatid"`
	Token  string `json:"token"`
}

type ConfigMiddlewares struct {
	GeoIP ConfigMiddlewaresGeoIP `json:"geoip"`
}

type ConfigMiddlewaresGeoIP struct {
	Enabled bool `json:"enabled" default:"false"`
}

type ConfigFilters struct {
	IPFilter ConfigFiltersIPFilter `json:"ipfilters"`
}

type ConfigFiltersIPFilter struct {
	Enabled bool     `json:"enabled"`
	Cidrs   []string `json:"cidrs"`
}
