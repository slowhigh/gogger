package sub

import (
	"github.com/Slowhigh/gogger/consumer/internal/entity/proto"
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

func (c Controller) ConsumeAccessLog(accessLogPb *proto.AccessLog) bool {
	return c.msgUsecase.CreateAccessLog(accessLogPb)
}