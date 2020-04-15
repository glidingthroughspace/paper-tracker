package router

import (
	"sync"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	httpParamIDName        = "id"
	httpParamTemplIDName   = "templid"
	httpParamExecIDName    = "execid"
	httpQueryDirectionName = "direction"
)

type HttpRouter struct {
	engine *gin.Engine
}

type RouterMode string

const (
	DebugMode   RouterMode = "debug"
	ReleaseMode RouterMode = "release"
)

// SetMode configures the operational mode of the router(s)
// Currently this just sets gin's mode to reduce spammy logs in production
func SetMode(mode RouterMode) {
	switch mode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
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
	r.engine.Use(static.Serve("/", static.LocalFile("static", true)))

	r.buildAppRoomAPIRoutes()
	r.buildAppTrackerAPIRoutes()
	r.buildAppTemplateAPIRoutes()
	r.buildAppExecAPIRoutes()
	r.buildAppExportAPIRoutes()
}
