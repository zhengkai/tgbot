package tgbot

import (
	"fmt"
	"strconv"
	"strings"
)

type Bot struct {
	ID    uint64
	Token string
}

func NewBot(token string) (*Bot, error) {

	b := &Bot{
		Token: token,
	}

	id := strings.Split(token, `:`)[0]
	b.ID, _ = strconv.ParseUint(id, 10, 64)
	if b.ID == 0 {
		return nil, fmt.Errorf(`invalid token: %s`, token)
	}

	return b, nil
}
