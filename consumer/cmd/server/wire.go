//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Slowhigh/gogger/consumer/infra"
	"github.com/Slowhigh/gogger/consumer/infra/consumer"
	"github.com/Slowhigh/gogger/consumer/internal/adapter"
	"github.com/Slowhigh/gogger/consumer/internal/usecase"
	"github.com/google/wire"
)

func InitServer() (*consumer.Consumer, error) {
	wire.Build(infra.InfraSet, adapter.AdapterSet, usecase.UsecaseSet)
	return &consumer.Consumer{}, nil
}
