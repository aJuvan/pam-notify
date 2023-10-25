package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aJuvan/pam-notify/config"
	"github.com/aJuvan/pam-notify/middleware"
)

type webhookBody struct {
	Username string              `json:"username"`
	Content  string              `json:"content"`
	Embeds   []webhookBodyEmbeds `json:"embeds"`
}

type webhookBodyEmbeds struct {
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Color       int                       `json:"color"`
	Fields      []webhookBodyEmbedsFields `json:"fields"`
	Footer      webhookBodyEmbedsFooter   `json:"footer"`
}

type webhookBodyEmbedsFields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type webhookBodyEmbedsFooter struct {
	Text string `json:"text"`
}

func Notify(userData config.UserData, notifier config.ConfigNotifier, middlewareData middleware.MiddlewareData) error {
	body := webhookBody{
		Username: "PAM-Notify",
		Content:  "",
		Embeds: []webhookBodyEmbeds{
			{
				Title:       "User login",
				Description: "New user login received",
				Color:       0xffc000,
				Fields: []webhookBodyEmbedsFields{
					{
						Name:   "Server",
						Value:  userData.Server,
						Inline: true,
					},
					{
						Name:   "Service",
						Value:  userData.Service,
						Inline: true,
					},
					{
						Name:   "Username",
						Value:  userData.Username,
						Inline: true,
					},
					{
						Name:   "Remote host",
						Value:  userData.Rhost,
						Inline: true,
					},
				},
				Footer: webhookBodyEmbedsFooter{
					Text: "PAM-Notify",
				},
			},
		},
	}

	if middlewareData.GeoIP != nil {
		body.Embeds = append(body.Embeds, webhookBodyEmbeds{
			Title:       "GeoIP",
			Description: "GeoIP data",
			Color:       0xaaaa00,
			Fields: []webhookBodyEmbedsFields{
				{
					Name:   "Continent",
					Value:  middlewareData.GeoIP.Continent,
					Inline: true,
				},
				{
					Name:   "Country",
					Value:  middlewareData.GeoIP.Country,
					Inline: true,
				},
				{
					Name:   "City",
					Value:  middlewareData.GeoIP.City,
					Inline: true,
				},
				{
					Name:   "GeoLocation",
					Value:  fmt.Sprintf("%f,%f", middlewareData.GeoIP.Latitude, middlewareData.GeoIP.Longitude),
					Inline: true,
				},
				{
					Name:   "Hostname",
					Value:  middlewareData.GeoIP.Hostname,
					Inline: true,
				},
			},
		})
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		notifier.Url,
		"application/json",
		bytes.NewBuffer(bodyJson),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Discord returned %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
