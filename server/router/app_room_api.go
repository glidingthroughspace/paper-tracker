package router

import (
	"net/http"
	"paper-tracker/managers"
	"paper-tracker/models"
	"paper-tracker/models/communication"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppRoomAPIRoutes() {
	room := r.engine.Group("/room")
	room.GET("", r.roomListHandler())
	room.POST("", r.roomCreateHandler())
	room.PUT("/:id", extractID(), r.roomUpdateHandler())
	room.DELETE("/:id", extractID(), r.roomDeleteHandler())
}

func (r *HttpRouter) roomListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rooms, err := managers.GetRoomManager().GetAllRooms()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, rooms)
	}
}

func (r *HttpRouter) roomCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		room := &models.Room{}
		err := ctx.BindJSON(room)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to room")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetRoomManager().CreateRoom(room)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) roomUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roomID := ctx.GetInt(httpParamIDName)

		room := &models.Room{}
		err := ctx.BindJSON(room)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to room")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		room.ID = roomID

		err = managers.GetRoomManager().UpdateRoom(room)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) roomDeleteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roomID := ctx.GetInt(httpParamIDName)

		err := managers.GetRoomManager().DeleteRoom(roomID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	}
}
