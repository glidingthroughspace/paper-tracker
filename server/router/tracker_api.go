package router

import (
	"paper-tracker/managers"
	"paper-tracker/models/communication"

	coap "github.com/go-ocf/go-coap"
	log "github.com/sirupsen/logrus"
	"github.com/ugorji/go/codec"
)

func (r *CoapRouter) buildTrackerAPIRoutes() {
	r.addRoute("/tracker/new", &routeHandlers{Post: r.trackerNewHandler()})
	r.addRoute("/tracker/poll", &routeHandlers{Get: r.trackerPollHandler()})
	r.addRoute("/tracker/tracking", &routeHandlers{Post: r.trackerTrackingData()})
}

func (r *CoapRouter) trackerNewHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		reqLogger := log.WithField("clientIP", req.Client.RemoteAddr().String)

		tracker, err := managers.GetTrackerManager().NotifyNewTracker()
		if err != nil {
			r.writeError(w, coap.InternalServerError, err)
			reqLogger.WithField("err", err).Warning("Coap router: Failed to notify new tracker")
			return
		}

		r.writeCBOR(w, coap.Created, tracker)
	}
}

func (r *CoapRouter) trackerPollHandler() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		reqLogger := log.WithField("clientIP", req.Client.RemoteAddr().String)

		trackerID, err := r.extractTrackerID(req)
		if err != nil {
			r.writeError(w, coap.BadRequest, err)
			reqLogger.WithField("err", err).Warning("Coap router: Failed to extract tracker ID")
			return
		}

		reqLogger = reqLogger.WithField("trackerID", trackerID)

		cmd, err := managers.GetTrackerManager().PollCommand(trackerID)
		if err != nil {
			reqLogger.WithField("err", err).Warning("Coap router: Failed to poll command")
			r.writeError(w, coap.InternalServerError, err)
			return
		}

		r.writeCBOR(w, coap.Content, cmd)
	}
}

func (r *CoapRouter) trackerTrackingData() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		reqLogger := log.WithField("clientIP", req.Client.RemoteAddr().String)

		trackerID, err := r.extractTrackerID(req)
		if err != nil {
			r.writeError(w, coap.BadRequest, err)
			log.WithFields(log.Fields{"clientIP": req.Client.RemoteAddr().String, "err": err}).Warning("Coap router: Failed to extract tracker ID")
			return
		}

		reqLogger = reqLogger.WithField("trackerID", trackerID)

		dec := codec.NewDecoderBytes(req.Msg.Payload(), r.cborHandle)
		defer dec.Release()

		resp := &communication.TrackingCmdResponse{}
		err = dec.Decode(resp)
		if err != nil {
			reqLogger.WithField("err", err).Warning("Coap router: Failed decode tracking data")
		}

		err = managers.GetLearningManager().NewTrackingData(trackerID, resp.ScanResults)
		if err != nil {
			r.writeError(w, coap.InternalServerError, err)
			reqLogger.WithField("err", err).Warning("Coap router: Failed to save tracking data")
			return
		}

		w.SetCode(coap.Empty)
	}
}
