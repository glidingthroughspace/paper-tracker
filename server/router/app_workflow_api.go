package router

import (
	"net/http"
	"paper-tracker/managers"
	"paper-tracker/models"
	"paper-tracker/models/communication"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppWorkflowAPIRoutes() {
	workflow := r.engine.Group("/workflow")
	workflow.GET("", r.workflowListHandler())
	workflow.POST("", r.workflowCreateHandler())
	workflow.POST("/:id/start", extractID(), r.workflowCreateStartHandler())
	workflow.POST("/:id/step", extractID(), r.workflowCreateStepHandler())
}

func (r *HttpRouter) workflowListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workflows, err := managers.GetWorkflowManager().GetAllWorkflows()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowList request failed")
			return
		}
		ctx.JSON(http.StatusOK, workflows)
	}
}

func (r *HttpRouter) workflowCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workflow := &models.Workflow{}
		err := ctx.BindJSON(workflow)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to workflow")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().CreateWorkflow(workflow)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowCreate request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowCreateStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workflowID := models.WorkflowID(ctx.GetInt(httpParamIDName))

		step := &models.Step{}
		err := ctx.BindJSON(step)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to step")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().CreateWorkflowStart(workflowID, step)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowCreateStart request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowCreateStepHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stepRequest := &communication.CreateStepRequest{}
		err := ctx.BindJSON(stepRequest)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to step")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().AddStep(stepRequest.PrevStepID, stepRequest.DecisionLabel, stepRequest.Step)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowCreateStep request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}
