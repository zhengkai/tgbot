package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	jp "github.com/buger/jsonparser"
	"github.com/zhengkai/tgbot/pb"

	"google.golang.org/protobuf/proto"
)

var ErrHTTPCode = errors.New(`http code is not 200`)
var ErrHTTPBody = errors.New(`read body fail`)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func httpGetJSON(url string, d any) error {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set(`Content-Type`, `application/json`)
	client := &http.Client{
		Timeout: 70 * time.Second,
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

	// fmt.Println(url, len(ab))

	ok, err := jp.GetBoolean(ab, `ok`)
	if !ok {
		// fmt.Println(string(ab))
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

func httpPostJSON(url string, d any, re proto.Message) error {

	ab, err := json.Marshal(d)
	if err != nil {
		return err
	}
	// fmt.Println(`send:`, string(ab))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(ab))
	if err != nil {
		return err
	}
	req.Header.Set(`Content-Type`, `application/json`)

	rsp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	ab, err = io.ReadAll(rsp.Body)
	if err != nil {

		return ErrHTTPBody
	}

	if rsp.StatusCode != http.StatusOK {
		fmt.Println(`http code not ok`, rsp.StatusCode)
		o := &pb.Error{}
		e2 := json.Unmarshal(ab, o)
		if e2 != nil {
			return ErrHTTPCode
		}

		// fmt.Println(rsp.StatusCode)
		return errors.New(o.Description)
	}

	if re == nil {
		return nil
	}

	return json.Unmarshal(ab, re)
}

func httpFetch(url string, w io.Writer) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf(`failed to initiate request: %v`, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(`bad status: %s`, resp.Status)
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return fmt.Errorf(`failed to copy response body: %v`, err)
	}

	return nil
}
