package adapter

import (
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/rest"
	"github.com/google/wire"
)

var AdapterSet = wire.NewSet(rest.NewController)
