package router

import (
	"paper-tracker/managers"
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type HttpRouter struct {
	engine     *gin.Engine
	trackerMgr *managers.TrackerManager
	roomMgr    *managers.RoomManager
}

func NewHttpRouter(trackerMgr *managers.TrackerManager, roomMgr *managers.RoomManager) *HttpRouter {
	r := &HttpRouter{
		engine:     gin.New(),
		trackerMgr: trackerMgr,
		roomMgr:    roomMgr,
	}
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
