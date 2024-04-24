package adapter

import (
	"github.com/Slowhigh/gogger/consumer/internal/adapter/controller"
	"github.com/google/wire"
)

var AdapterSet = wire.NewSet(controller.NewController)
