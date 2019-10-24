package router

import "github.com/gin-gonic/gin"

type Router struct {
	Router *gin.Engine
}

func NewRouter() *Router {
	r := &Router{
		Router: gin.New(),
	}
	r.buildRoutes()
	return r
}

func (r *Router) buildRoutes() {
	r.buildTrackerAPIRoutes()
}
