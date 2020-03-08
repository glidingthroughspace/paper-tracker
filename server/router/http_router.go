package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	httpParamIDName        = "id"
	httpParamTemplIDName   = "tempid"
	httpParamExecIDName    = "execid"
	httpQueryDirectionName = "direction"
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
	r.buildAppRoomAPIRoutes()
	r.buildAppTrackerAPIRoutes()
	r.buildAppTemplateAPIRoutes()
	r.buildAppExecAPIRoutes()
}
