package adapter

import (
	"github.com/Slowhigh/gogger/consumer/internal/adapter/controller/sub"
	"github.com/google/wire"
)

var AdapterSet = wire.NewSet(sub.NewController)
