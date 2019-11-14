package router

import (
	"paper-tracker/managers"
	"paper-tracker/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *HttpRouter) buildAppAPIRoutes() {
	r.engine.GET("/tracker", r.trackerListHandler())
}

func (r *HttpRouter) trackerListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackers, err := managers.GetTrackerManager().GetAllTrackers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, trackers)
	}
}
