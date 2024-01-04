package infra

import (
	"github.com/Slowhigh/gogger/producer/infra/producer"
	"github.com/Slowhigh/gogger/producer/infra/router"
	"github.com/google/wire"
)

var InfraSet = wire.NewSet(producer.NewProducer, router.NewRouter)
