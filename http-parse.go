package tgbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	jp "github.com/buger/jsonparser"
	"github.com/zhengkai/tgbot/pb"
)

func parseHTTPRsp(rsp *http.Response, re any) error {
	defer rsp.Body.Close()

	ab, err := io.ReadAll(rsp.Body)
	if err != nil {
		return ErrHTTPBody
	}

	err = parseHTTPResult(ab, re)
	if rsp.StatusCode != http.StatusOK {
		if err == ErrNoOK {
			return ErrHTTPCode
		}
		return err
	}
	return err
}

func parseHTTPResult(ab []byte, re any) error {

	fmt.Println(`parseHTTPResult`, string(ab))

	ok, _ := jp.GetBoolean(ab, `ok`)
	if !ok {
		o := &pb.Error{}
		json.Unmarshal(ab, o)
		if o.Description != `` {
			return errors.New(o.Description)
		}
		return ErrNoOK
	}

	ab, _, _, _ = jp.Get(ab, `result`)
	if len(ab) == 0 {
		return errors.New(`api fail: no result`)
	}

	return json.Unmarshal(ab, re)
}
