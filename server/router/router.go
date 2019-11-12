package router

import (
	"encoding/json"
	"paper-tracker/models"

	coap "github.com/go-ocf/go-coap"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	Mux *coap.ServeMux
}

func NewRouter() *Router {
	r := &Router{
		Mux: coap.NewServeMux(),
	}
	r.buildRoutes()
	return r
}

func (r *Router) buildRoutes() {
	r.buildTrackerAPIRoutes()
}

type routeHandlers struct {
	Get    coap.HandlerFunc
	Post   coap.HandlerFunc
	Put    coap.HandlerFunc
	Delete coap.HandlerFunc
}

func (r *Router) addRoute(path string, handlers *routeHandlers) {
	r.Mux.Handle(path, coap.HandlerFunc(func(w coap.ResponseWriter, req *coap.Request) {
		reqType := req.Msg.Code()
		if reqType == coap.GET && handlers.Get != nil {
			handlers.Get.ServeCOAP(w, req)
		} else if reqType == coap.POST && handlers.Post != nil {
			handlers.Post.ServeCOAP(w, req)
		} else if reqType == coap.PUT && handlers.Put != nil {
			handlers.Put.ServeCOAP(w, req)
		} else if reqType == coap.DELETE && handlers.Delete != nil {
			handlers.Delete.ServeCOAP(w, req)
		} else {
			r.writeJSON(w, coap.MethodNotAllowed, &models.ErrorResponse{Error: "Method not allowed"})
		}
	}))
}

func (r *Router) writeJSON(w coap.ResponseWriter, status coap.COAPCode, body interface{}) (err error) {
	w.SetContentFormat(coap.AppJSON)
	w.SetCode(status)
	data, err := json.Marshal(body)
	if err != nil {
		log.Error("Failed to marshal body to json: %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Error("Failed to write JSON response: %v", err)
		return
	}
	return
}

func (r *Router) writeError(w coap.ResponseWriter, err error) error {
	return r.writeJSON(w, coap.InternalServerError, &models.ErrorResponse{Error: err.Error()})
}
