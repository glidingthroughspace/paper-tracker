package router

import (
	"net/http"
	"paper-tracker/models/communication"
	"strconv"

	"github.com/gin-gonic/gin"
)

func extractID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Set(ginID, id)
	}
}
