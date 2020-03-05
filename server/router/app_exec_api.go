package router

import (
	"net/http"
	"paper-tracker/managers"
	"paper-tracker/models"
	"paper-tracker/models/communication"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppExecAPIRoutes() {
	workflow := r.engine.Group("/workflow")

	exec := workflow.Group("/exec")
	exec.GET("", r.workflowExecListHandler())
	exec.POST("", r.workflowExecStartHandler())
	exec.POST("/:execID/progress/:id", extractID("execID", httpParamExecIDName), extractSimpleID(), r.workflowExecProgressHandler())
	exec.POST("/:execID/cancel", extractID("execID", httpParamExecIDName), r.workflowExecCancelHandler())
}

func (r *HttpRouter) workflowExecListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		execs, err := managers.GetWorkflowExecManager().GetAllExec()
		if err != nil {
			log.WithError(err).Warn("Failed to get all workflow execs")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, execs)
	}
}

func (r *HttpRouter) workflowExecStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		exec := &models.WorkflowExec{}
		err := ctx.BindJSON(exec)
		if err != nil {
			log.WithError(err).Error("Failed to unmarshal json to workflow execution")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowExecManager().StartExecution(exec)
		if err != nil {
			log.WithFields(log.Fields{"exec": exec, "err": err}).Warn("Failed to start workflow execution")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, exec)
	}
}

func (r *HttpRouter) workflowExecProgressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		execID := models.WorkflowExecID(ctx.GetInt(httpParamExecIDName))
		stepID := models.StepID(ctx.GetInt(httpParamIDName))

		err := managers.GetWorkflowExecManager().ProgressToStep(execID, stepID)
		if err != nil {
			log.WithFields(log.Fields{"execID": execID, "stepID": stepID, "err": err}).Warn("Failed to progres exec to step")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowExecCancelHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		execID := models.WorkflowExecID(ctx.GetInt(httpParamExecIDName))

		err := managers.GetWorkflowExecManager().CancelExec(execID)
		if err != nil {
			log.WithFields(log.Fields{"execID": execID, "err": err}).Warn("Failed to cancel exec")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	}
}
