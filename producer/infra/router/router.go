package router

import (
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/http"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(ctrl http.Controller) Router {
	r := gin.Default()
	r.POST("/log/access", func(c *gin.Context) { accessLogHandler(c, ctrl) })
	return Router{router: r}
}

func (r Router) Run() error {
	return r.router.Run(":5000")
}
