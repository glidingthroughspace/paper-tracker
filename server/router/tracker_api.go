package router

import (
	"paper-tracker/managers"
	"paper-tracker/models"
	"strconv"

	coap "github.com/go-ocf/go-coap"
)

func (r *CoapRouter) buildTrackerAPIRoutes() {
	r.addRoute("/tracker/notify-new", &routeHandlers{Post: r.trackerNotifyHandler()})
	r.addRoute("/tracker/poll", &routeHandlers{Get: r.trackerPollHandler()})
}

func (r *CoapRouter) trackerNotifyHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		tracker, err := managers.GetTrackerManager().NotifyNewTracker()
		if err != nil {
			r.writeError(w, err)
			return
		}

		r.writeCBOR(w, coap.Created, tracker)
	}
}

func (r *CoapRouter) trackerPollHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		params := r.parseQuery(req)
		trackerIDStr, ok := params["trackerid"]
		if !(ok && trackerIDStr != nil) {
			r.writeCBOR(w, coap.BadRequest, &models.ErrorResponse{Error: "trackerid not found in query"})
			return
		}
		trackerID, err := strconv.Atoi(*trackerIDStr)
		if err != nil {
			r.writeCBOR(w, coap.BadRequest, &models.ErrorResponse{Error: "trackerid is not an integer"})
		}

		cmd, err := managers.GetTrackerManager().PollCommand(trackerID)
		if err != nil {
			r.writeError(w, err)
		}

		r.writeCBOR(w, coap.Content, cmd)
	}
}
