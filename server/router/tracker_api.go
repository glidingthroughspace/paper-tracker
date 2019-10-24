package router

import (
	"net/http"
	"paper-tracker/managers"
	"paper-tracker/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) buildTrackerAPIRoutes() {
	trackerAPI := r.Router.Group("/api/tracker")
	trackerAPI.POST("/notify-new", r.trackerNotifyHandler())
	trackerAPI.GET("/:id/poll", r.trackerPollHandler())
}

func (r *Router) trackerNotifyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracker, err := managers.GetTrackerManager().NotifyNewTracker()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, tracker)
	}
}

func (r *Router) trackerPollHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}
