package tgbot

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"google.golang.org/protobuf/proto"
)

func httpUpload(url string, m map[string]any, re proto.Message) error {

	body, contentType, err := httpUploadBuildBody(m)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf(`Error creating request: %s`, err)
	}
	req.Header.Set(`Content-Type`, contentType)

	client := &http.Client{
		Timeout: 70 * time.Second,
	}
	rsp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf(`Error sending request: %s`, err)
	}

	return parseHTTPRsp(rsp, re)
}

func httpUploadBuildBody(m map[string]any) (io.Reader, string, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range m {
		switch rv := v.(type) {
		case string:
			writer.WriteField(k, rv)
		case *os.File:
			part, err := writer.CreateFormFile(k, rv.Name())
			if err != nil {
				return nil, ``, err
			}
			_, err = io.Copy(part, rv)
			if err != nil {
				return nil, ``, err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, ``, fmt.Errorf(`Error closing writer: %s`, err)
	}

	return body, writer.FormDataContentType(), nil
}
