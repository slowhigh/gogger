package http

import (
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/http/dto"
	"github.com/Slowhigh/gogger/producer/internal/entity"
	"github.com/Slowhigh/gogger/producer/internal/usecase/message"
)

type Controller struct {
	msgUsecase message.MessageUsecase
}

func NewController(msgUsecase message.MessageUsecase) Controller {
	return Controller{
		msgUsecase: msgUsecase,
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
