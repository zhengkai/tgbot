package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

var ErrHTTPCode = errors.New(`http code is not 200`)
var ErrHTTPBody = errors.New(`read body fail`)
var ErrNoOK = errors.New(`api fail: no "ok:true"`)

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
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	parseHTTPRsp(rsp, d)

	// fmt.Println(url, len(ab))

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

	return parseHTTPRsp(rsp, re)
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
