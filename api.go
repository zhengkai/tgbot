package tgbot

import (
	"encoding/json"
)

type api struct {
	bot     *Bot
	urlBase string
}

func (b *Bot) API() *api {
	return &api{
		bot:     b,
		urlBase: `https://api.telegram.org/bot` + b.Token + `/`,
	}
}

func (a *api) GetMe() (*Me, error) {
	ab, err := httpGet(a.urlBase + `getMe`)
	if err != nil {
		return nil, err
	}

	d := &Me{}
	err = json.Unmarshal(ab, d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
