package router

import (
	"fmt"
	"paper-tracker/managers"
	"paper-tracker/models/communication"

	coap "github.com/go-ocf/go-coap"
	"github.com/ugorji/go/codec"
)

func (r *CoapRouter) buildTrackerAPIRoutes() {
	r.addRoute("/tracker/new", &routeHandlers{Post: r.trackerNewHandler()})
	r.addRoute("/tracker/poll", &routeHandlers{Get: r.trackerPollHandler()})
	r.addRoute("/tracker/tracking", &routeHandlers{Post: r.trackerTrackingData()})
}

func (r *CoapRouter) trackerNewHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		tracker, err := managers.GetTrackerManager().NotifyNewTracker()
		if err != nil {
			r.writeError(w, coap.InternalServerError, err)
			return
		}

		r.writeCBOR(w, coap.Created, tracker)
	}
}

func (r *CoapRouter) trackerPollHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		trackerID, err := r.extractTrackerID(req)
		if err != nil {
			r.writeError(w, coap.BadRequest, err)
			return
		}

		cmd, err := managers.GetTrackerManager().PollCommand(trackerID)
		if err != nil {
			r.writeError(w, coap.InternalServerError, err)
			return
		}

		r.writeCBOR(w, coap.Content, cmd)
	}
}

func (r *CoapRouter) trackerTrackingData() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		trackerID, err := r.extractTrackerID(req)
		if err != nil {
			r.writeError(w, coap.BadRequest, err)
			return
		}

		dec := codec.NewDecoderBytes(req.Msg.Payload(), r.cborHandle)
		defer dec.Release()

		resp := &communication.TrackingCmdResponse{}
		err = dec.Decode(resp)
		if err != nil {
			fmt.Println(err)
		}

		err = managers.GetLearningManager().NewTrackingData(trackerID, resp.ScanResults)
		if err != nil {
			r.writeError(w, coap.InternalServerError, err)
			return
		}

		w.SetCode(coap.Empty)
	}
}
