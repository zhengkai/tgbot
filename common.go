package tgbot

import "github.com/zhengkai/tgbot/pb"

type ISendPhoto interface {
	GetReplyParameters() *pb.ReplyParameters
}

func parsePhotoMap(m map[string]any, arg ISendPhoto) {
	if a := arg.GetReplyParameters(); a != nil {
		m[`reply_parameters`] = jsonStr(a)
	}
}
