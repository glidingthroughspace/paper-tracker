package router

import (
	"net/http"
	"paper-tracker/models/communication"
	"strconv"

	"github.com/gin-gonic/gin"
)

func extractSimpleID() gin.HandlerFunc {
	return extractID("id", httpParamIDName)
}

func extractTemplID() gin.HandlerFunc {
	return extractID("templID", httpParamTemplIDName)
}

func extractID(paramName, contextName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param(paramName))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Set(contextName, id)
	}
}
