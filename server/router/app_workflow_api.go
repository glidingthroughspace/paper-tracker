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

	template := workflow.Group("/template")
	template.GET("", r.workflowTemplateListHandler())
	template.POST("", r.workflowTemplateCreateHandler())
	template.POST("/:id/start", extractID(), r.workflowTemplateCreateStartHandler())
	template.POST("/:id/step", extractID(), r.workflowTemplateCreateStepHandler())

	exec := workflow.Group("/exec")
	exec.GET("", r.workflowExecListHandler())
	exec.POST("/start", r.workflowExecStartHandler())
}

func (r *HttpRouter) workflowTemplateListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workflows, err := managers.GetWorkflowManager().GetAllTemplates()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowList request failed")
			return
		}
		ctx.JSON(http.StatusOK, workflows)
	}
}

func (r *HttpRouter) workflowTemplateCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workflow := &models.WorkflowTemplate{}
		err := ctx.BindJSON(workflow)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to workflow")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().CreateTemplate(workflow)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowCreate request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowTemplateCreateStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workflowID := models.WorkflowTemplateID(ctx.GetInt(httpParamIDName))

		step := &models.Step{}
		err := ctx.BindJSON(step)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to step")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().CreateTemplateStart(workflowID, step)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowCreateStart request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowTemplateCreateStepHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stepRequest := &communication.CreateStepRequest{}
		err := ctx.BindJSON(stepRequest)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to step")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().AddTemplateStep(stepRequest.PrevStepID, stepRequest.DecisionLabel, stepRequest.Step)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowCreateStep request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowExecListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		execs, err := managers.GetWorkflowManager().GetAllExec()
		if err != nil {
			log.WithField("err", err).Warn("Failed to get all workflow execs")
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
			log.WithField("err", err).Error("Failed to unmarshal json to workflow execution")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().StartExecution(exec)
		if err != nil {
			log.WithFields(log.Fields{"exec": exec, "err": err}).Warn("Failed to start workflow execution")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	}
}
