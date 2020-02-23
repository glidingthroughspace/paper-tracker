package managers

import (
	"errors"
	"paper-tracker/models"
	"paper-tracker/repositories"
	"time"

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

func (mgr *WorkflowManager) GetTemplate(templateID models.WorkflowTemplateID) (workflow *models.WorkflowTemplate, err error) {
	getWorkflowLog := log.WithField("workflowID", templateID)

	workflow, err = mgr.workflowRep.GetTemplateByID(templateID)
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

func (mgr *WorkflowManager) GetAllExec() (execs []*models.WorkflowExec, err error) {
	rawExecs, err := mgr.workflowRep.GetAllExec()
	if err != nil {
		log.WithField("err", err).Error("Failed to get all execs")
		return
	}

	execs = make([]*models.WorkflowExec, len(rawExecs))
	for it, rawExec := range rawExecs {
		execs[it], err = mgr.GetExec(rawExec.ID)
		if err != nil {
			log.WithFields(log.Fields{"err": err, "rawID": rawExec.ID}).Error("Failed to get workflow exec for list")
			continue
		}
	}

	return
}

func (mgr *WorkflowManager) GetExec(execID models.WorkflowExecID) (exec *models.WorkflowExec, err error) {
	execLog := log.WithField("execID", execID)

	exec, err = mgr.workflowRep.GetExecByID(execID)
	if err != nil {
		execLog.WithField("err", err).Error("Failed to get workflow exec")
		return
	}

	infos, err := mgr.workflowRep.GetExecStepInfoForExecID(exec.ID)
	if err != nil {
		execLog.WithField("err", err).Error("Failed to get infos for exec")
		return
	}

	exec.StepInfos = make(map[models.StepID]*models.ExecStepInfo, len(infos))
	for _, info := range infos {
		exec.StepInfos[info.StepID] = info
	}

	return
}

func (mgr *WorkflowManager) StartExecution(exec *models.WorkflowExec) (err error) {
	startExecLog := log.WithField("exec", exec)

	tracker, err := GetTrackerManager().GetTrackerByID(exec.TrackerID)
	if err != nil {
		startExecLog.WithField("err", err).Warn("Failed to get tracker for starting execution")
		return
	}

	if tracker.Status != models.StatusIdle {
		startExecLog.Warn("Tracker not in idle for starting execution")
		err = errors.New("tracker not in idle mode")
		return
	}

	template, err := mgr.GetTemplate(exec.TemplateID)
	if err != nil {
		startExecLog.WithField("err", err).Warn("Failed to get template for starting execution")
		return
	}

	if template.StartStep == models.StepID(0) {
		startExecLog.Warn("Workflow template does not have any steps for starting execution")
		err = errors.New("template does not have any steps")
		return
	}

	exec.Completed = false
	exec.CurrentStepID = template.StartStep
	exec.StartedOn = time.Now()
	err = mgr.workflowRep.CreateExec(exec)
	if err != nil {
		startExecLog.WithField("err", err).Warn("Failed to create workflow exec")
		return
	}

	err = GetTrackerManager().SetTrackerStatus(exec.TrackerID, models.StatusTracking)
	if err != nil {
		startExecLog.WithField("err", err).Error("Failed to set tracker to status tracking - error ignored for now")
		err = nil
	}

	startInfo := &models.ExecStepInfo{
		ExecID:    exec.ID,
		StepID:    template.StartStep,
		StartedOn: time.Now(),
	}
	err = mgr.workflowRep.CreateExecStepInfo(startInfo)
	if err != nil {
		startExecLog.WithFields(log.Fields{"info": startInfo, "err": err}).Error("Failed to create info for start step - error ignored for now")
		err = nil
	}

	for stepID, info := range exec.StepInfos {
		if stepID == template.StartStep {
			startInfo.Decision = info.Decision
			err = mgr.workflowRep.UpdateExecStepInfo(startInfo)
			if err != nil {
				startExecLog.WithFields(log.Fields{"info": startInfo, "err": err}).Error("Failed to update info for start step - error ignored for now")
				err = nil
			}
			continue
		}
		info.ExecID = exec.ID
		info.StepID = stepID
		err = mgr.workflowRep.CreateExecStepInfo(info)
		if err != nil {
			startExecLog.WithFields(log.Fields{"info": info, "err": err}).Error("Failed to create exec info - error ignored for now")
			err = nil
		}
	}

	return
}
