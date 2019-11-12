package router

import (
	"paper-tracker/managers"

	coap "github.com/go-ocf/go-coap"
)

func (r *Router) buildTrackerAPIRoutes() {
	r.addRoute("/tracker/notify-new", &routeHandlers{Post: r.trackerNotifyHandler()})
	r.addRoute("/tracker/:id/poll", &routeHandlers{Get: r.trackerPollHandler()})
}

func (r *Router) trackerNotifyHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		tracker, err := managers.GetTrackerManager().NotifyNewTracker()
		if err != nil {
			r.writeError(w, err)
			return
		}

		r.writeJSON(w, coap.Created, tracker)
	}
}

func (r *Router) trackerPollHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		//TODO
	}
}
