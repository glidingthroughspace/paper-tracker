package router

import (
	"paper-tracker/managers"
	"paper-tracker/models/communication"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *HttpRouter) buildAppAPIRoutes() {
	r.engine.GET("/tracker", r.trackerListHandler())
	r.engine.POST("/tracker/:id/learn/start", r.trackerLearnStartHandler())
	r.engine.GET("/tracker/:id/learn/status", r.trackerLearnStatusHandler())

	r.engine.GET("/room", r.roomListHandler())
}

func (r *HttpRouter) trackerListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackers, err := managers.GetTrackerManager().GetAllTrackers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, trackers)
	}
}

func (r *HttpRouter) trackerLearnStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		learnTime, err := managers.GetLearningManager().StartLearning(trackerID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &communication.LearningStartResponse{LearnTimeSec: learnTime})
	}
}

func (r *HttpRouter) trackerLearnStatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		done, ssids, err := managers.GetLearningManager().GetLearningStatus(trackerID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &communication.LearningStatusResponse{Done: done, SSIDs: ssids})
	}
}

func (r *HttpRouter) roomListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rooms, err := managers.GetRoomManager().GetAllRooms()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, rooms)
	}
}
