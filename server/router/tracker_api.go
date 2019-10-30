package router

import (
	"paper-tracker/managers"

	coap "github.com/go-ocf/go-coap"
)

func (r *Router) buildTrackerAPIRoutes() {
	r.Mux.Handle("/notify-new", r.trackerNotifyHandler())
	r.Mux.Handle("/:id/poll", r.trackerPollHandler())
}

func (r *Router) trackerNotifyHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		_, err := managers.GetTrackerManager().NotifyNewTracker()
		if err != nil {
			//c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		//c.JSON(http.StatusOK, tracker)
	}
}

func (r *Router) trackerPollHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		//TODO
	}
}
