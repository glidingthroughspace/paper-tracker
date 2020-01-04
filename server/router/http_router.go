package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	httpParamIDName = "id"
)

type HttpRouter struct {
	engine *gin.Engine
}

func NewHttpRouter() *HttpRouter {
	r := &HttpRouter{
		engine: gin.New(),
	}
	r.engine.Use(gin.Logger())
	r.buildRoutes()
	return r
}

func (r *HttpRouter) Serve(addr string, wg *sync.WaitGroup) {
	log.Error(r.engine.Run(addr))
	wg.Done()
}

func (r *HttpRouter) buildRoutes() {
	r.buildAppAPIRoutes()
}
