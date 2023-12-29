package http

import (
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/http/dto"
	"github.com/Slowhigh/gogger/producer/internal/entity"
	"github.com/Slowhigh/gogger/producer/internal/entity/interactor"
	"github.com/Slowhigh/gogger/producer/internal/usecase"
)

type Controller struct {
	msgUsecase usecase.MessageUsecase
}

func NewController(mp interactor.MessageProducer) Controller {
	return Controller{
		msgUsecase: usecase.NewMessageUsecase(mp),
	}
}

func (c Controller) ProduceAccessLog(dto dto.AccessLogDto) bool {
	accessLog := entity.AccessLog{
		Timestamp:    dto.Timestamp,
		IsNormalMode: *dto.IsNormalMode,
		IsLogin:      *dto.IsLogin,
		UserName:     dto.UserName,
		DeviceName:   dto.DeviceName,
		Ip:           dto.Ip,
	}

	return c.msgUsecase.ProduceAccessLog(accessLog)
}
