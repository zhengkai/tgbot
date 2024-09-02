package tgbot

import (
	"fmt"
	"io"
	"net/url"

	"github.com/zhengkai/tgbot/pb"
)

// https://core.telegram.org/bots/api#getfile
func (a *api) GetFile(fileID string, w io.Writer) error {

	u := fmt.Sprintf(`%sgetFile?file_id=%s`, a.urlBase, url.QueryEscape(fileID))

	d := &pb.File{}
	err := httpGetJSON(u, d)
	if err != nil {
		return err
	}

	u = fmt.Sprintf(`https://api.telegram.org/file/bot%s/%s`, a.bot.Token, d.FilePath)
	return httpFetch(u, w)
}
