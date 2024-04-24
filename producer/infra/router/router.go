package router

import (
	"fmt"

	"github.com/Slowhigh/gogger/producer/infra/config"
	"github.com/Slowhigh/gogger/producer/infra/router/handler"
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/rest"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	conf   *config.Config
}

func NewRouter(conf *config.Config, ctrl rest.Controller) Router {
	if conf.Http.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.POST("/log/access", func(c *gin.Context) { handler.AccessLogHandler(c, ctrl) })
	return Router{
		router: r,
		conf:   conf,
	}
}

func (r Router) Run() error {
	return r.router.Run(fmt.Sprintf(":%d", r.conf.Http.Port))
}
