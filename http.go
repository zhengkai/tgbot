package tgbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	jp "github.com/buger/jsonparser"
)

func httpGetJSON(url string, d any) error {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set(`Content-Type`, `application/json`)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	ab, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	fmt.Println(string(ab))

	ok, err := jp.GetBoolean(ab, `ok`)
	if !ok {
		return errors.New(`api fail: no "ok:true"`)
	}

	ab, t, _, err := jp.Get(ab, `result`)
	if len(ab) == 0 {
		return errors.New(`api fail: no result`)
	}
	if t != jp.Object && t != jp.Array {
		return errors.New(`api fail: result is not object`)
	}

	err = json.Unmarshal(ab, d)
	if err != nil {
		return err
	}

	return nil
}
