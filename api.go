package tgbot

import (
	"fmt"
	"os"

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

func (a *api) GetUpdates(offset uint64) ([]*pb.Update, error) {

	url := a.urlBase + `getUpdates`

	if offset > 0 {
		url += fmt.Sprintf(`?offset=%d&timeout=60`, offset)
	}

	d := []*pb.Update{}
	err := httpGetJSON(url, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (a *api) SendChatAction(m *pb.SendChatAction) error {
	url := a.urlBase + `sendChatAction`
	var ok bool
	return httpPostJSON(url, m, &ok)
}

func (a *api) SendMessage(m any) (*pb.Message, error) {
	url := a.urlBase + `sendMessage`

	re := &pb.Message{}

	err := httpPostJSON(url, m, re)

	return nil, err
}

func (a *api) SendPhoto(chatId string, file *os.File, arg *pb.SendPhoto) (*pb.PhotoReturn, error) {
	url := a.urlBase + `sendPhoto`

	m := map[string]any{
		`chat_id`: chatId,
		`photo`:   file,
	}
	parsePhotoMap(m, arg)

	o := &pb.PhotoReturn{}
	err := httpUpload(url, m, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (a *api) SendAnimation(chatId string, file *os.File, arg *pb.SendAnimation) (*pb.PhotoReturn, error) {
	url := a.urlBase + `sendAnimation`

	m := map[string]any{
		`chat_id`:   chatId,
		`animation`: file,
	}
	parsePhotoMap(m, arg)

	o := &pb.PhotoReturn{}
	err := httpUpload(url, m, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}
