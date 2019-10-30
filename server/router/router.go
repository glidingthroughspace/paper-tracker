package router

import coap "github.com/go-ocf/go-coap"

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
