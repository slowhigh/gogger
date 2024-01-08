package infra

import (
	"github.com/Slowhigh/gogger/consumer/infra/config"
	"github.com/Slowhigh/gogger/consumer/infra/consumer"
	"github.com/Slowhigh/gogger/consumer/infra/database"
	"github.com/Slowhigh/gogger/consumer/infra/database/repository"
	"github.com/google/wire"
)

var InfraSet = wire.NewSet(
	config.NewConfig,
	consumer.NewConsumer,
	database.NewDatabase,
	repository.NewAccessLog,
)
