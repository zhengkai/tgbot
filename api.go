package tgbot

import (
	"fmt"

	"github.com/zhengkai/tgbot/pb"
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

func (a *api) GetMe() (*pb.User, error) {

	d := &pb.User{}
	err := httpGetJSON(a.urlBase+`getMe`, d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (a *api) GetUpdates(offset int) ([]*pb.Update, error) {

	url := a.urlBase + `getUpdates`

	if offset > 0 {
		url += fmt.Sprintf(`?offset=%d`, offset)
	}

	d := []*pb.Update{}
	err := httpGetJSON(url, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
