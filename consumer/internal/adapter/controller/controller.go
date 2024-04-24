package controller

import (
	"github.com/Slowhigh/gogger/consumer/internal/adapter/controller/dto"
	"github.com/Slowhigh/gogger/consumer/internal/entity"
	"github.com/Slowhigh/gogger/consumer/internal/usecase/message"
)

type Controller struct {
	msgUsecase message.MessageUsecase
}

func NewController(mu message.MessageUsecase) Controller {
	return Controller{
		msgUsecase: mu,
	}
}

func (c Controller) ConsumeAccessLog(dto dto.AccessLogDto) bool {
	return c.msgUsecase.CreateAccessLog(entity.AccessLog{
		Timestamp:    dto.Timestamp,
		IsNormalMode: dto.IsNormalMode,
		IsLogin:      dto.IsLogin,
		UserName:     dto.UserName,
		DeviceName:   dto.DeviceName,
		Ip:           dto.Ip,
	})
}
