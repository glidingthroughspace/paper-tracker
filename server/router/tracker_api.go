package router

import (
	"paper-tracker/managers"

	coap "github.com/go-ocf/go-coap"
)

func (r *CoapRouter) buildTrackerAPIRoutes() {
	r.addRoute("/tracker/notify-new", &routeHandlers{Post: r.trackerNotifyHandler()})
	r.addRoute("/tracker/:id/poll", &routeHandlers{Get: r.trackerPollHandler()})
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
		//TODO
	}
}
