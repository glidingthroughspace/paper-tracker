package router

import (
	"paper-tracker/managers"
	"paper-tracker/models"
	"paper-tracker/models/communication"

	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppTrackerAPIRoutes() {
	tracker := r.engine.Group("/tracker")
	tracker.GET("", r.trackerListHandler())
	tracker.PUT("/:id", extractID(), r.trackerUpdateHandler())
	tracker.DELETE("/:id", extractID(), r.trackerDeleteHandler())

	trackerLearn := tracker.Group("/:id/learn", extractID())
	trackerLearn.POST("/start", r.trackerLearnStartHandler())
	trackerLearn.GET("/status", r.trackerLearnStatusHandler())
	trackerLearn.POST("/finish", r.trackerLearnFinishHandler())

}

func (r *HttpRouter) trackerListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackers, err := managers.GetTrackerManager().GetAllTrackers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("TrackerList request failed")
			return
		}
		ctx.JSON(http.StatusOK, trackers)
	}
}

func (r *HttpRouter) trackerUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID := models.TrackerID(ctx.GetInt(httpParamIDName))

		tracker := &models.Tracker{}
		err := ctx.BindJSON(tracker)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to tracker")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		tracker, err = managers.GetTrackerManager().UpdateTracker(trackerID, tracker.Label)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("TrackerUpdate request failed")
			return
		}
		ctx.JSON(http.StatusOK, tracker)
	}
}

func (r *HttpRouter) trackerDeleteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID := models.TrackerID(ctx.GetInt(httpParamIDName))

		err := managers.GetTrackerManager().DeleteTracker(trackerID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("TrackerDelete request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) trackerLearnStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID := models.TrackerID(ctx.GetInt(httpParamIDName))

		learnTime, err := managers.GetLearningManager().StartLearning(trackerID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("TrackerLearnStart request failed")
			return
		}
		ctx.JSON(http.StatusOK, &communication.LearningStartResponse{LearnTimeSec: learnTime})
	}
}

func (r *HttpRouter) trackerLearnStatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID := models.TrackerID(ctx.GetInt(httpParamIDName))

		done, ssids, err := managers.GetLearningManager().GetLearningStatus(trackerID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("TrackerLearnStatus request failed")
			return
		}
		ctx.JSON(http.StatusOK, &communication.LearningStatusResponse{Done: done, SSIDs: ssids})
	}
}

func (r *HttpRouter) trackerLearnFinishHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackerID := models.TrackerID(ctx.GetInt(httpParamIDName))

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
			log.WithField("err", err).Warn("TrackerLearnFinish request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}
