package router

import (
	"net/http"
	"paper-tracker/managers"
	"paper-tracker/models/communication"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppExportAPIRoutes() {
	r.engine.GET("/export.xlsx", r.exportHandler())
}

func (r *HttpRouter) exportHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := managers.GetExportManager().GenerateExport(ctx.Writer)
		if err != nil {
			log.WithError(err).Error("Failed to generate export")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
	}
}
