package router

import (
	"errors"
	"paper-tracker/models"
	"paper-tracker/models/communication"
	"strconv"
	"strings"
	"sync"

	coap "github.com/go-ocf/go-coap"
	log "github.com/sirupsen/logrus"
	"github.com/ugorji/go/codec"
)

type CoapRouter struct {
	mux        *coap.ServeMux
	cborHandle codec.Handle
}

func NewCoapRouter() *CoapRouter {
	r := &CoapRouter{
		mux:        coap.NewServeMux(),
		cborHandle: &codec.CborHandle{},
	}
	r.mux.DefaultHandleFunc(r.notfound())
	r.buildRoutes()
	return r
}

func (r *CoapRouter) Serve(network, addr string, wg *sync.WaitGroup) {
	log.Error(coap.ListenAndServe("udp", ":5688", r.mux))
	wg.Done()
}

func (r *CoapRouter) buildRoutes() {
	r.buildTrackerAPIRoutes()
}

type routeHandlers struct {
	Get    coap.HandlerFunc
	Post   coap.HandlerFunc
	Put    coap.HandlerFunc
	Delete coap.HandlerFunc
}

func (r *CoapRouter) addRoute(path string, handlers *routeHandlers) {
	r.mux.Handle(path, r.loggingMiddleware(r.methodSwitchMiddleware(handlers)))
}

func (r *CoapRouter) writeCBOR(w coap.ResponseWriter, status coap.COAPCode, body interface{}) (err error) {
	w.SetContentFormat(coap.AppCBOR)
	w.SetCode(status)

	enc := codec.NewEncoder(w, r.cborHandle)
	defer enc.Release()
	err = enc.Encode(body)
	if err != nil {
		log.WithError(err).Error("Failed to write or encode CBOR response")
		return
	}
	return
}

func (r *CoapRouter) writeError(w coap.ResponseWriter, code coap.COAPCode, err error) error {
	return r.writeCBOR(w, code, &communication.ErrorResponse{Error: err.Error()})
}

func (r *CoapRouter) parseQuery(req *coap.Request) (paramMap map[string]*string) {
	paramMap = make(map[string]*string)

	queryParams := req.Msg.Query()
	for _, param := range queryParams {
		paramSplit := strings.SplitN(param, "=", 2)
		if len(paramSplit) == 2 {
			paramMap[paramSplit[0]] = &paramSplit[1]
		} else {
			paramMap[paramSplit[0]] = nil
		}
	}

	return
}

func (r *CoapRouter) extractTrackerID(req *coap.Request) (trackerID models.TrackerID, err error) {
	params := r.parseQuery(req)
	trackerIDStr, ok := params["trackerid"]
	if !(ok && trackerIDStr != nil) {
		err = errors.New("trackerid not found in query")
		return
	}
	trackerIDInt, err := strconv.Atoi(*trackerIDStr)
	if err != nil {
		err = errors.New("trackerid is not an integer")
		return
	}
	trackerID = models.TrackerID(trackerIDInt)

	return
}
