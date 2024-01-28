package message

import (
	"log/slog"

	"github.com/Slowhigh/gogger/consumer/internal/entity"
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

func (mu MessageUsecase) CreateAccessLog(accessLog entity.AccessLog) bool {
	if err := mu.accessLogRepo.Create(&accessLog); err != nil {
		slog.Error(err.Error())
		return false
	}

	return true
}
