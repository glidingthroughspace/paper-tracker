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
	template.GET("/:tempID/step/:id", extractID(), r.workflowTemplateGetStepHandler())
	template.POST("/:id/revision", extractID(), r.workflowTemplateNewRevisionHandler())

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
		template := &models.WorkflowTemplate{}
		err := ctx.BindJSON(template)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to workflow template")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().CreateTemplate(template)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowTemplateCreate request failed")
			return
		}
		ctx.JSON(http.StatusOK, template)
	}
}

func (r *HttpRouter) workflowTemplateCreateStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamIDName))

		step := &models.Step{}
		err := ctx.BindJSON(step)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to step")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowManager().CreateTemplateStart(templateID, step)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowTemplateCreateStart request failed")
			return
		}
		ctx.JSON(http.StatusOK, step)
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
			log.WithField("err", err).Warn("WorkflowTemplateCreateStep request failed")
			return
		}
		ctx.JSON(http.StatusOK, stepRequest.Step)
	}
}

func (r *HttpRouter) workflowTemplateGetStepHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stepID := models.StepID(ctx.GetInt(httpParamIDName))

		step, err := managers.GetWorkflowManager().GetStepByID(stepID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithField("err", err).Warn("WorkflowTemplateGetStep request failed")
			return
		}
		ctx.JSON(http.StatusOK, step)
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
		ctx.JSON(http.StatusOK, exec)
	}
}

func (r *HttpRouter) workflowTemplateNewRevisionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamIDName))

		revisionRequest := &communication.CreateRevisionRequest{}
		err := ctx.BindJSON(revisionRequest)
		if err != nil {
			log.WithField("err", err).Error("Failed to unmarshal json to create revision request")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		newTemplate, err := managers.GetWorkflowManager().CreateNewRevision(templateID, revisionRequest.RevisionLabel)
		if err != nil {
			log.WithFields(log.Fields{"templateID": templateID, "err": err}).Warn("Failed to create new template revision")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
		}
		ctx.JSON(http.StatusOK, newTemplate)
	}
}
