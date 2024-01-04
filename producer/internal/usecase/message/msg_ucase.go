package message

import (
	"log/slog"

	"github.com/Slowhigh/gogger/producer/internal/entity"
	"github.com/Slowhigh/gogger/producer/internal/entity/interactor"
	"github.com/Slowhigh/gogger/producer/internal/entity/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MessageUsecase struct {
	messageProducer interactor.Producer
}

func NewMessageUsecase(mp interactor.Producer) MessageUsecase {
	return MessageUsecase{messageProducer: mp}
}

func (mu MessageUsecase) ProduceAccessLog(msg entity.AccessLog) bool {
	accessLogProto := proto.AccessLog{
		Timestamp:    timestamppb.New(msg.Timestamp),
		IsNormalMode: msg.IsNormalMode,
		IsLogin:      msg.IsLogin,
		UserName:     msg.UserName,
		DeviceName:   msg.DeviceName,
		Ip:           msg.Ip,
	}

	err := mu.messageProducer.Produce(&accessLogProto)
	if err != nil {
		slog.Error(err.Error())
		return false
	}

	return true
}
