package managers

import (
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

var workflowManager *WorkflowManager

type WorkflowManager struct {
	workflowRep repositories.WorkflowRepository
}

func CreateWorkflowManager(workflowRep repositories.WorkflowRepository) *WorkflowManager {
	if workflowManager != nil {
		return workflowManager
	}

	workflowManager = &WorkflowManager{
		workflowRep: workflowRep,
	}

	return workflowManager
}

func GetWorkflowManager() *WorkflowManager {
	return workflowManager
}

func (mgr *WorkflowManager) CreateWorkflow(workflow *models.Workflow) (err error) {
	workflow.ID = 0
	err = mgr.workflowRep.CreateWorkflow(workflow)
	if err != nil {
		log.WithFields(log.Fields{"workflow": workflow, "err": err}).Error("Failed to create workflow")
		return
	}
	return
}

func (mgr *WorkflowManager) CreateWorkflowStart(workflowID models.WorkflowID, step *models.Step) (err error) {
	workflowStartLog := log.WithFields(log.Fields{"workflowID": workflowID, "step": step})

	step.ID = 0
	err = mgr.workflowRep.CreateStep(step)
	if err != nil {
		workflowStartLog.WithField("err", err).Error("Failed to create step")
		return
	}

	workflow, err := mgr.workflowRep.GetWorkflowByID(workflowID)
	if err != nil {
		workflowStartLog.WithField("err", err).Error("Failed to get workflow to create start")
		mgr.workflowRep.DeleteStep(step.ID)
		return
	}

	workflow.StartStep = step.ID
	err = mgr.workflowRep.UpdateWorkflow(workflow)
	if err != nil {
		workflowStartLog.WithField("err", err).Error("Failed to update workflow to create start")
		mgr.workflowRep.DeleteStep(step.ID)
		return
	}
	return
}
