package message

import (
	"log/slog"

	"github.com/Slowhigh/gogger/producer/internal/entity"
	"github.com/Slowhigh/gogger/producer/internal/entity/proto"
	"github.com/Slowhigh/gogger/producer/internal/usecase/interactor"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MessageUsecase struct {
	producer interactor.Producer
}

func NewMessageUsecase(p interactor.Producer) MessageUsecase {
	return MessageUsecase{
		producer: p,
	}
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

	if err := mu.producer.ProduceAccessLog(&accessLogProto); err != nil {
		slog.Error(err.Error())
		return false
	}

	return true
}
