package message

import (
	"log/slog"

	"github.com/Slowhigh/gogger/consumer/internal/entity"
	"github.com/Slowhigh/gogger/consumer/internal/entity/proto"
	"github.com/Slowhigh/gogger/consumer/internal/usecase/interactor"
)

type MessageUsecase struct {
	accessLogRepo interactor.AccessLog
}

func NewMessageUsecase(accessLogRepo interactor.AccessLog) MessageUsecase {
	return MessageUsecase{
		accessLogRepo: accessLogRepo,
	}
}

func (mu MessageUsecase) CreateAccessLog(accessLogPb *proto.AccessLog) bool {
	accessLog := entity.AccessLog{
		Timestamp:    accessLogPb.Timestamp.AsTime(),
		IsNormalMode: accessLogPb.IsNormalMode,
		IsLogin:      accessLogPb.IsLogin,
		UserName:     accessLogPb.UserName,
		DeviceName:   accessLogPb.DeviceName,
		Ip:           accessLogPb.Ip,
	}

	if err := mu.accessLogRepo.Create(&accessLog); err != nil {
		slog.Error(err.Error())
		return false
	}

	return true
}
