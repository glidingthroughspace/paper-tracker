package router

import (
	"paper-tracker/models/communication"
	"strings"

	coap "github.com/go-ocf/go-coap"
	log "github.com/sirupsen/logrus"
)

func (r *CoapRouter) notfound() coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		log.WithFields(log.Fields{
			"path":          "/" + strings.Join(req.Msg.Path(), "/"),
			"code":          req.Msg.Code().String(),
			"src":           req.Client.RemoteAddr(),
			"contentFormat": req.Msg.Option(coap.ContentFormat),
		}).Info("Path not found")
		w.SetCode(coap.NotFound)
		w.Write([]byte{})
	}
}

func (r *CoapRouter) loggingMiddleware(next coap.HandlerFunc) coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
		log.WithFields(log.Fields{
			"path":          "/" + strings.Join(req.Msg.Path(), "/"),
			"code":          req.Msg.Code().String(),
			"src":           req.Client.RemoteAddr(),
			"contentFormat": req.Msg.Option(coap.ContentFormat),
		}).Info("Handle request")

		next.ServeCOAP(w, req)
	}
}

func (r *CoapRouter) methodSwitchMiddleware(handlers *routeHandlers) coap.HandlerFunc {
	return func(w coap.ResponseWriter, req *coap.Request) {
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
			r.writeCBOR(w, coap.MethodNotAllowed, &communication.ErrorResponse{Error: "Method not allowed"})
		}
	}
}
