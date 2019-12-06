package router

import (
	"paper-tracker/managers"
	"paper-tracker/models"
	"paper-tracker/models/communication"

	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppAPIRoutes() {
	tracker := r.engine.Group("/tracker")
	tracker.GET("/", r.trackerListHandler())

	trackerLearn := tracker.Group("/:id/learn", extractID())
	trackerLearn.POST("/start", r.trackerLearnStartHandler())
	trackerLearn.GET("/status", r.trackerLearnStatusHandler())
	trackerLearn.POST("/finish", r.trackerLearnFinishHandler())

	r.engine.GET("/room", r.roomListHandler())
	r.engine.POST("/room", r.roomCreateHandler())
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
		trackerID := ctx.GetInt(httpParamIDName)

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
		trackerID := ctx.GetInt(httpParamIDName)

		done, ssids, err := managers.GetLearningManager().GetLearningStatus(trackerID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &communication.LearningStatusResponse{Done: done, SSIDs: ssids})
	}
}

func (r *HttpRouter) trackerLearnFinishHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID := ctx.GetInt(httpParamIDName)

		req := &communication.LearningFinishRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to learn finish request")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetLearningManager().FinishLearning(trackerID, req.RoomID, req.SSIDs)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
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

func (r *HttpRouter) roomCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		room := &models.Room{}
		err := ctx.BindJSON(room)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to room")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetRoomManager().CreateRoom(room)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	}
}
