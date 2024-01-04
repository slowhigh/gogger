package usecase

import (
	"github.com/Slowhigh/gogger/producer/internal/usecase/message"
	"github.com/google/wire"
)

var UsecaseSet = wire.NewSet(message.NewMessageUsecase)
