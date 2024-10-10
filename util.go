package tgbot

import "encoding/json"

func jsonStr(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
