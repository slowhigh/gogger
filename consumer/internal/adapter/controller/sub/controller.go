package sub

import (
	"github.com/Slowhigh/gogger/consumer/internal/adapter/controller/sub/dto"
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
	accessLog := entity.AccessLog{
		Timestamp:    dto.Timestamp,
		IsNormalMode: dto.IsNormalMode,
		IsLogin:      dto.IsLogin,
		UserName:     dto.UserName,
		DeviceName:   dto.DeviceName,
		Ip:           dto.Ip,
	}
	return c.msgUsecase.CreateAccessLog(accessLog)
}
