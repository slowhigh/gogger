package usecase

import (
	"github.com/Slowhigh/gogger/consumer/internal/usecase/message"
	"github.com/google/wire"
)

var UsecaseSet = wire.NewSet(message.NewMessageUsecase)
