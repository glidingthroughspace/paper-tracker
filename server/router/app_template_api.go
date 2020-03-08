package router

import (
	"net/http"
	"paper-tracker/managers"
	"paper-tracker/models"
	"paper-tracker/models/communication"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *HttpRouter) buildAppTemplateAPIRoutes() {
	workflow := r.engine.Group("/workflow")

	template := workflow.Group("/template")
	template.GET("", r.workflowTemplateListHandler())
	template.POST("", r.workflowTemplateCreateHandler())
	template.DELETE("/:templID", extractTemplID(), r.workflowTemplateDeleteHandler())
	template.POST("/:templID/start", extractTemplID(), r.workflowTemplateCreateStartHandler())
	template.POST("/:templID/step", extractTemplID(), r.workflowTemplateCreateStepHandler())
	template.GET("/:templID/step/:id", extractTemplID(), extractSimpleID(), r.workflowTemplateGetStepHandler())
	template.PUT("/:templID/step/:id", extractTemplID(), extractSimpleID(), r.workflowTemplateUpdateStepHandler())
	template.DELETE("/:templID/step/:id", extractTemplID(), extractSimpleID(), r.workflowTemplateDeleteStepHandler())
	template.POST("/:templID/step/:id/move", extractTemplID(), extractSimpleID(), r.workflowTemplateMoveStepHandler())
	template.POST("/:templID/revision", extractTemplID(), r.workflowTemplateNewRevisionHandler())
}

func (r *HttpRouter) workflowTemplateListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workflows, err := managers.GetWorkflowTemplateManager().GetAllTemplates()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowList request failed")
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
			log.WithError(err).Error("Failed to unmarshal json to workflow template")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowTemplateManager().CreateTemplate(template)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowTemplateCreate request failed")
			return
		}
		ctx.JSON(http.StatusOK, template)
	}
}

func (r *HttpRouter) workflowTemplateDeleteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamTemplIDName))

		err := managers.GetWorkflowTemplateManager().DeleteTemplate(templateID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowTemplateDelete request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowTemplateCreateStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamTemplIDName))

		step := &models.Step{}
		err := ctx.BindJSON(step)
		if err != nil {
			log.WithError(err).Error("Failed to unmarshal json to step")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowTemplateManager().CreateTemplateStart(templateID, step)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowTemplateCreateStart request failed")
			return
		}
		ctx.JSON(http.StatusOK, step)
	}
}

func (r *HttpRouter) workflowTemplateCreateStepHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamTemplIDName))

		stepRequest := &communication.CreateStepRequest{}
		err := ctx.BindJSON(stepRequest)
		if err != nil {
			log.WithError(err).Error("Failed to unmarshal json to step")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		err = managers.GetWorkflowTemplateManager().AddTemplateStep(templateID, stepRequest.PrevStepID, stepRequest.DecisionLabel, stepRequest.Step)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowTemplateCreateStep request failed")
			return
		}
		ctx.JSON(http.StatusOK, stepRequest.Step)
	}
}

func (r *HttpRouter) workflowTemplateGetStepHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamTemplIDName))
		stepID := models.StepID(ctx.GetInt(httpParamIDName))

		step, err := managers.GetWorkflowTemplateManager().GetStepByID(templateID, stepID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowTemplateGetStep request failed")
			return
		}
		ctx.JSON(http.StatusOK, step)
	}
}

func (r *HttpRouter) workflowTemplateUpdateStepHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamTemplIDName))
		stepID := models.StepID(ctx.GetInt(httpParamIDName))

		step := &models.Step{}
		err := ctx.BindJSON(step)
		if err != nil {
			log.WithError(err).Error("Failed to unmarshal json to step")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		step.ID = stepID

		err = managers.GetWorkflowTemplateManager().UpdateStep(templateID, step)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowTemplateUpdateStep request failed")
			return
		}
		ctx.JSON(http.StatusOK, step)
	}
}

func (r *HttpRouter) workflowTemplateDeleteStepHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamTemplIDName))
		stepID := models.StepID(ctx.GetInt(httpParamIDName))

		err := managers.GetWorkflowTemplateManager().DeleteStep(templateID, stepID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowTemplateUpdateStep request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowTemplateMoveStepHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamTemplIDName))
		stepID := models.StepID(ctx.GetInt(httpParamIDName))
		direction := communication.StepMoveDirectionFromString(ctx.Query(httpQueryDirectionName))

		err := managers.GetWorkflowTemplateManager().MoveStep(templateID, stepID, direction)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			log.WithError(err).Warn("WorkflowTemplateMoveStep request failed")
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (r *HttpRouter) workflowTemplateNewRevisionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		templateID := models.WorkflowTemplateID(ctx.GetInt(httpParamTemplIDName))

		revisionRequest := &communication.CreateRevisionRequest{}
		err := ctx.BindJSON(revisionRequest)
		if err != nil {
			log.WithError(err).Error("Failed to unmarshal json to create revision request")
			ctx.JSON(http.StatusBadRequest, &communication.ErrorResponse{Error: err.Error()})
			return
		}

		newTemplate, err := managers.GetWorkflowTemplateManager().CreateNewRevision(templateID, revisionRequest.RevisionLabel)
		if err != nil {
			log.WithFields(log.Fields{"templateID": templateID, "err": err}).Warn("Failed to create new template revision")
			ctx.JSON(http.StatusInternalServerError, &communication.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, newTemplate)
	}
}
