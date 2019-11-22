package router

import (
	"paper-tracker/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *HttpRouter) buildAppAPIRoutes() {
	r.engine.GET("/tracker", r.trackerListHandler())
	r.engine.POST("/tracker/:id/learn/start", r.trackerLearnStartHandler())
}

func (r *HttpRouter) trackerListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackers, err := r.trackerMgr.GetAllTrackers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, trackers)
	}
}

func (r *HttpRouter) trackerLearnStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{Error: err.Error()})
			return
		}

		learnTime, err := r.trackerMgr.StartLearning(trackerID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{Error: err.Error()})
			return
		}
		//TODO: Move to a model
		ctx.JSON(http.StatusOK, struct{ LearnTimeSec int }{learnTime})
	}
}
