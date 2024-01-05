//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Slowhigh/gogger/producer/infra"
	"github.com/Slowhigh/gogger/producer/infra/router"
	"github.com/Slowhigh/gogger/producer/internal/adapter"
	"github.com/Slowhigh/gogger/producer/internal/usecase"
	"github.com/google/wire"
)

func InitServer() (router.Router, error) {
	wire.Build(infra.InfraSet, adapter.AdapterSet, usecase.UsecaseSet)
	return router.Router{}, nil
}
