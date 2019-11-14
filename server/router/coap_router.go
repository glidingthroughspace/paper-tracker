package router

import (
	"paper-tracker/models"
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
		log.Error("Failed to write or encode CBOR response: %v", err)
		return
	}
	return
}

func (r *CoapRouter) writeError(w coap.ResponseWriter, err error) error {
	return r.writeCBOR(w, coap.InternalServerError, &models.ErrorResponse{Error: err.Error()})
}
