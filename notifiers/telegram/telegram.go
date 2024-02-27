package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/middleware"
)

type webhookBody struct {
	ChatId         string `json:"chat_id"`
	Text           string `json:"text"`
	ProtectContent bool   `json:"protect_content"`
}

func Run(notifier *config.ConfigNotifierTelegram, user *config.UserData, middlewareData *middleware.MiddlewareData) error {
	contentArray := []string{
		"**PAM Notification**",
		"Server: " + user.Server,
		"Service: " + user.Service,
		"User: " + user.Username,
		"Remote host: " + user.Rhost,
	}

	if middlewareData.GeoIP != nil {
		tmp := []string{
			"",
			"Continent: " + middlewareData.GeoIP.Continent,
			"Country: " + middlewareData.GeoIP.Country,
			"City: " + middlewareData.GeoIP.City,
			"Geolocation: " + fmt.Sprintf("%f, %f", middlewareData.GeoIP.Latitude, middlewareData.GeoIP.Longitude),
			"Hostname: " + middlewareData.GeoIP.Hostname,
		}
		contentArray = append(contentArray, tmp...)
	}

	content := strings.Join(contentArray, "\n")
	bodyJson, err := json.Marshal(webhookBody{
		ChatId:         notifier.ChatId,
		Text:           content,
		ProtectContent: true,
	})
	if err != nil {
		return err
	}

	logger := config.Logger.With().Str("module", "notifiers").Str("notifier", "telegram").Logger()
	logger.Debug().Str("chat_id", notifier.ChatId).Str("token", notifier.Token).Msg("Sending message to Telegram")

	resp, err := http.Post(
		"https://api.telegram.org/bot"+notifier.Token+"/sendMessage",
		"application/json",
		bytes.NewBuffer(bodyJson),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Telegram API returned status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
