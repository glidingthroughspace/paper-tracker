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

func (mgr *WorkflowManager) CreateTemplate(workflow *models.WorkflowTemplate) (err error) {
	workflow.ID = 0
	err = mgr.workflowRep.CreateTemplate(workflow)
	if err != nil {
		log.WithFields(log.Fields{"workflow": workflow, "err": err}).Error("Failed to create workflow")
		return
	}
	return
}

func (mgr *WorkflowManager) CreateTemplateStart(workflowID models.WorkflowTemplateID, step *models.Step) (err error) {
	workflowStartLog := log.WithFields(log.Fields{"workflowID": workflowID, "step": step})

	step.ID = 0
	err = mgr.workflowRep.CreateStep(step)
	if err != nil {
		workflowStartLog.WithField("err", err).Error("Failed to create step")
		return
	}

	workflow, err := mgr.workflowRep.GetTemplateByID(workflowID)
	if err != nil {
		workflowStartLog.WithField("err", err).Error("Failed to get workflow to create start")
		mgr.workflowRep.DeleteStep(step.ID)
		return
	}

	workflow.StartStep = step.ID
	err = mgr.workflowRep.UpdateTemplate(workflow)
	if err != nil {
		workflowStartLog.WithField("err", err).Error("Failed to update workflow to create start")
		mgr.workflowRep.DeleteStep(step.ID)
		return
	}
	return
}

// TODO: Fix adding in between to steps
func (mgr *WorkflowManager) AddTemplateStep(prevStepID models.StepID, decisionLabel string, step *models.Step) (err error) {
	addStepLog := log.WithFields(log.Fields{"prevStepID": prevStepID, "step": step})

	step.ID = 0
	err = mgr.workflowRep.CreateStep(step)
	if err != nil {
		addStepLog.WithField("err", err).Error("Failed to create step to add step")
		return
	}

	_, err = mgr.workflowRep.GetStepByID(prevStepID)
	if mgr.workflowRep.IsRecordNotFoundError(err) {
		addStepLog.WithField("err", err).Warn("Previous step not found to add step")
		mgr.workflowRep.DeleteStep(step.ID)
		return
	} else if err != nil {
		addStepLog.WithField("err", err).Error("Failed to get previous step to add step")
		mgr.workflowRep.DeleteStep(step.ID)
		return
	}

	nextStep := &models.NextStep{
		PrevID:        prevStepID,
		NextID:        step.ID,
		DecisionLabel: decisionLabel,
	}
	err = mgr.workflowRep.CreateNextStep(nextStep)
	if err != nil {
		addStepLog.WithField("err", err).Error("Failed to get insert nextStep to add step")
		mgr.workflowRep.DeleteStep(step.ID)
		return
	}

	return
}

func (mgr *WorkflowManager) GetAllTemplates() (workflows []*models.WorkflowTemplate, err error) {
	rawWorkflows, err := mgr.workflowRep.GetAllTemplates()
	if err != nil {
		log.WithField("err", err).Error("Failed to get all raw workflows")
		return
	}

	workflows = make([]*models.WorkflowTemplate, len(rawWorkflows))
	for it, raw := range rawWorkflows {
		workflows[it], err = mgr.GetTemplate(raw.ID)
		if err != nil {
			log.WithFields(log.Fields{"err": err, "rawID": raw.ID}).Error("Failed to get workflow for list")
			continue
		}
	}
	return
}

func (mgr *WorkflowManager) GetTemplate(workflowID models.WorkflowTemplateID) (workflow *models.WorkflowTemplate, err error) {
	getWorkflowLog := log.WithField("workflowID", workflowID)

	workflow, err = mgr.workflowRep.GetTemplateByID(workflowID)
	if mgr.workflowRep.IsRecordNotFoundError(err) {
		getWorkflowLog.Warn("Workflow not found")
		return
	} else if err != nil {
		getWorkflowLog.WithField("err", err).Error("Failed to get workflow")
		return
	}

	workflow.Steps, err = mgr.getStepsFromStart(workflow.StartStep, getWorkflowLog)
	if err != nil {
		getWorkflowLog.WithField("err", err).Error("Failed to get steps")
		return
	}

	return
}

func (mgr *WorkflowManager) getStepsFromStart(startStepID models.StepID, getLog *log.Entry) (steps []*models.Step, err error) {
	getStepsFromStartLog := getLog.WithField("startStepID", startStepID)
	steps = make([]*models.Step, 0)
	currentStepID := startStepID

	for currentStepID > 0 {
		currentStep, err := mgr.workflowRep.GetStepByID(currentStepID)
		if mgr.workflowRep.IsRecordNotFoundError(err) {
			getStepsFromStartLog.Warn("StartStep not found")
			break
		} else if err != nil {
			getStepsFromStartLog.WithField("err", err).Error("Failed to get startStep")
			break
		}

		steps = append(steps, currentStep)

		currentStepID, err = mgr.workflowRep.GetLinearNextStepID(currentStepID)
		if mgr.workflowRep.IsRecordNotFoundError(err) {
			getStepsFromStartLog.Info("No next linear step")
		} else if err != nil {
			getStepsFromStartLog.WithField("err", err).Warn("Failed to get next linear step ID")
		}
	}

	for _, step := range steps {
		decisions, err := mgr.workflowRep.GetDecisions(step.ID)
		if err != nil {
			getStepsFromStartLog.WithFields(log.Fields{"err": err, "stepID": step.ID}).Warn("Failed to get decisions for step")
			continue
		}
		step.Options = make(map[string][]*models.Step)
		for _, decision := range decisions {
			step.Options[decision.DecisionLabel], err = mgr.getStepsFromStart(decision.NextID, getLog)
			if err != nil {
				getStepsFromStartLog.WithFields(log.Fields{"err": err, "decision": decision}).Error("Failed to get steps for decision")
				continue
			}
		}
	}

	return
}
