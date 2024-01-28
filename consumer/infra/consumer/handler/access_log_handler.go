package handler

import (
	"context"

	"github.com/Slowhigh/gogger/consumer/internal/adapter/controller/sub/dto"
	"github.com/Slowhigh/gogger/consumer/internal/entity/proto"
	"github.com/memphisdev/memphis.go"
	protobuf "google.golang.org/protobuf/proto"
)

func AccessLogHandler(consume func(dto dto.AccessLogDto) bool, msgs []*memphis.Msg, err error, ctx context.Context) {
	if err != nil {
		return
	}

	for _, msg := range msgs {
		var accessLogPb proto.AccessLog
		err = protobuf.Unmarshal(msg.Data(), &accessLogPb)
		if err != nil {
			panic(err)
		}

		dto := dto.AccessLogDto{
			Timestamp:    accessLogPb.Timestamp.AsTime(),
			IsNormalMode: accessLogPb.IsNormalMode,
			IsLogin:      accessLogPb.IsLogin,
			UserName:     accessLogPb.UserName,
			DeviceName:   accessLogPb.DeviceName,
			Ip:           accessLogPb.Ip,
		}

		if consume(dto) {
			msg.Ack()
		}
	}

}
