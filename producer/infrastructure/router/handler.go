package router

import (
	"log/slog"

	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/http"
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/http/dto"
	"github.com/gin-gonic/gin"
)

func accessLogHandler(c *gin.Context, ctrl http.Controller) {
	var dto dto.AccessLogDto

	if err := c.ShouldBind(&dto); err != nil {
		slog.Error(err.Error())
		c.Status(BAD_REQUEST)
		return
	}

	if !ctrl.ProduceAccessLog(dto) {
		c.Status(INTERNAL_SERVER_ERROR)
		return
	}

	c.Status(OK)
}
