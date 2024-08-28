package tgbot

import (
	"errors"
	"io"
	"net/http"

	jp "github.com/buger/jsonparser"
)

func httpGet(url string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(`Content-Type`, `application/json`)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	ab, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	ok, err := jp.GetBoolean(ab, `ok`)
	if !ok {
		return nil, errors.New(`api fail: no "ok:true"`)
	}

	ab, t, _, err := jp.Get(ab, `result`)
	if len(ab) == 0 {
		return nil, errors.New(`api fail: no result`)
	}
	if t != jp.Object {
		return nil, errors.New(`api fail: result is not object`)
	}

	return ab, nil
}
