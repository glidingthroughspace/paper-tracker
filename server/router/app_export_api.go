package router

import (
	"net/http"
	"paper-tracker/managers"
	"paper-tracker/models/communication"
	"paper-tracker/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppExportAPIRoutes() {
	r.engine.GET("/export.xlsx", r.exportHandler())

	// TODO: Remove
	r.engine.POST("/email", r.emailHandler())
}

func (r *HttpRouter) emailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.SendMail("Hello Test")
	}
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
