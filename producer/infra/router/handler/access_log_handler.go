package handler

import (
	"log/slog"
	"net/http"

	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/rest"
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/rest/dto"
	"github.com/gin-gonic/gin"
)

func AccessLogHandler(c *gin.Context, ctrl rest.Controller) {
	var dto dto.AccessLogDto

	if err := c.ShouldBind(&dto); err != nil {
		slog.Error(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	if !ctrl.ProduceAccessLog(dto) {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
