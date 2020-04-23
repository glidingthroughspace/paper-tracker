package router

import (
	"net/http"
	"paper-tracker/config"
	"paper-tracker/managers"
	"paper-tracker/models/communication"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppConfigAPIRoutes() {
	r.engine.GET("/export.xlsx", r.configExportHandler())
	r.engine.GET("/config", r.configGetHandler())
	r.engine.POST("/config", r.configSetHandler())
}

func (r *HttpRouter) configExportHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := managers.GetExportManager().GenerateExport(ctx.Writer)
		if err != nil {
			log.WithError(err).Error("Failed to generate export")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
	}
}

func (r *HttpRouter) configGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cfg := config.GetEditableConfig()
		ctx.JSON(http.StatusOK, cfg)
	}
}

func (r *HttpRouter) configSetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cfg := &config.EditableConfigs{}
		err := ctx.BindJSON(cfg)
		if err != nil {
			log.WithError(err).Error("Failed to unmarshal json to config")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = config.UpdateEditableConfig(cfg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("UpdateConfig request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}
