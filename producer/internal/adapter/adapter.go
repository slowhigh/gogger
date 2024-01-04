package adapter

import (
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/http"
	"github.com/google/wire"
)

var AdapterSet = wire.NewSet(http.NewController)
