package managers

import (
	"errors"
	"fmt"
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

func (mgr *WorkflowManager) CreateTemplate(template *models.WorkflowTemplate) (err error) {
	template.ID = 0
	err = mgr.workflowRep.CreateTemplate(template)
	if err != nil {
		log.WithFields(log.Fields{"template": template, "err": err}).Error("Failed to create template")
		return
	}
	return
}

func (mgr *WorkflowManager) CreateTemplateStart(templateID models.WorkflowTemplateID, step *models.Step) (err error) {
	workflowStartLog := log.WithFields(log.Fields{"templateID": templateID, "step": step})

	template, err := mgr.GetTemplate(templateID)
	if err != nil || template.EditingLocked {
		workflowStartLog.WithField("err", err).Warn("Editing of template locked")
		return errors.New("Editing of template locked")
	}

	step.ID = 0
	err = mgr.workflowRep.CreateStep(step)
	if err != nil {
		workflowStartLog.WithField("err", err).Error("Failed to create step")
		return
	}

	template.StartStep = step.ID
	err = mgr.workflowRep.UpdateTemplate(template)
	if err != nil {
		workflowStartLog.WithField("err", err).Error("Failed to update template to create start")
		mgr.workflowRep.DeleteStep(step.ID)
		return
	}
	return
}

// TODO: Fix adding in between to steps
func (mgr *WorkflowManager) AddTemplateStep(templateID models.WorkflowTemplateID, prevStepID models.StepID, decisionLabel string, step *models.Step) (err error) {
	addStepLog := log.WithFields(log.Fields{"prevStepID": prevStepID, "step": step})

	template, err := mgr.GetTemplate(templateID)
	if err != nil || template.EditingLocked {
		addStepLog.WithField("err", err).Warn("Editing of template locked")
		return errors.New("Editing of template locked")
	}

	step.ID = 0
	err = mgr.workflowRep.CreateStep(step)
	if err != nil {
		addStepLog.WithField("err", err).Error("Failed to create step to add step")
		return
	}

	_, err = mgr.GetStepByID(templateID, prevStepID)
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

	execs, err := mgr.workflowRep.GetExecsByTemplateID(templateID)
	if err != nil {
		getWorkflowLog.WithField("err", err).Error("Failed to get execs of template - ignore for now")
		execs = make([]*models.WorkflowExec, 0)
		err = nil
	}
	if len(execs) > 0 {
		workflow.EditingLocked = true
	} else {
		workflow.EditingLocked = false
	}

	workflow.Steps, err = mgr.getStepsFromStart(templateID, workflow.StartStep, getWorkflowLog)
	if err != nil {
		getWorkflowLog.WithField("err", err).Error("Failed to get steps")
		return
	}

	return
}

func (mgr *WorkflowManager) getStepsFromStart(templateID models.WorkflowTemplateID, startStepID models.StepID, getLog *log.Entry) (steps []*models.Step, err error) {
	getStepsFromStartLog := getLog.WithField("startStepID", startStepID)
	steps = make([]*models.Step, 0)
	currentStepID := startStepID

	for currentStepID > 0 {
		currentStep, err := mgr.GetStepByID(templateID, currentStepID)
		if err != nil {
			break
		}

		steps = append(steps, currentStep)

		currentStepID, err = mgr.workflowRep.GetLinearNextStepID(currentStepID)
		if mgr.workflowRep.IsRecordNotFoundError(err) {
			getStepsFromStartLog.Info("No next linear step")
			break
		} else if err != nil {
			getStepsFromStartLog.WithField("err", err).Warn("Failed to get next linear step ID")
			break
		}
	}

	return
}

func (mgr *WorkflowManager) GetStepByID(templateID models.WorkflowTemplateID, stepID models.StepID) (step *models.Step, err error) {
	getStepLog := log.WithField("stepID", stepID)

	step, err = mgr.workflowRep.GetStepByID(stepID)
	if err != nil {
		getStepLog.WithField("err", err).Warn("Failed to get step by ID")
		return
	}
	if mgr.workflowRep.IsRecordNotFoundError(err) {
		getStepLog.Warn("Step not found")
		return
	} else if err != nil {
		getStepLog.WithField("err", err).Error("Failed to get startStep")
		return
	}

	decisions, err := mgr.workflowRep.GetDecisions(step.ID)
	if err != nil {
		getStepLog.WithField("err", err).Warn("Failed to get decisions for step")
		return
	}
	step.Options = make(map[string][]*models.Step)
	for _, decision := range decisions {
		step.Options[decision.DecisionLabel], err = mgr.getStepsFromStart(templateID, decision.NextID, getStepLog)
		if err != nil {
			getStepLog.WithFields(log.Fields{"err": err, "decision": decision}).Error("Failed to get steps for decision")
			continue
		}
	}

	return
}

func (mgr *WorkflowManager) UpdateStep(templateID models.WorkflowTemplateID, step *models.Step) (err error) {
	updateLog := log.WithFields(log.Fields{"templateID": templateID, "step": step})

	template, err := mgr.GetTemplate(templateID)
	if err != nil || template.EditingLocked {
		updateLog.WithField("err", err).Warn("Editing of template locked")
		return errors.New("Editing of template locked")
	}

	err = mgr.workflowRep.UpdateStep(step)
	if err != nil {
		updateLog.WithField("err", err).Error("Failed to update step")
		return
	}

	for decision, steps := range step.Options {
		nextStep := &models.NextStep{
			PrevID:        step.ID,
			NextID:        steps[0].ID,
			DecisionLabel: decision,
		}
		err = mgr.workflowRep.UpdateNextStep(nextStep)
		if err != nil {
			updateLog.WithField("err", err).Error("Failed to update decision")
			continue
		}
	}
	return
}

func (mgr *WorkflowManager) DeleteStep(templateID models.WorkflowTemplateID, stepID models.StepID) (err error) {
	deleteLog := log.WithFields(log.Fields{"templateID": templateID, "stepID": stepID})

	template, err := mgr.GetTemplate(templateID)
	if err != nil || template.EditingLocked {
		deleteLog.WithField("err", err).Warn("Editing of template locked")
		return errors.New("Editing of template locked")
	}

	step, err := mgr.GetStepByID(templateID, stepID)
	if err != nil {
		deleteLog.WithField("err", err).Error("Failed to get to be deleted step")
		return
	} else if len(step.Options) > 0 {
		deleteLog.Warn("Cannot delete step that has options")
		return errors.New("Cannot delete step that has options")
	}

	var fromStep *models.NextStep
	if template.StartStep != stepID {
		fromStep, err = mgr.workflowRep.GetNextStepByNextID(stepID)
		if err != nil {
			deleteLog.WithField("err", err).Error("Failed to get nextStep that points to be deleted step")
			return
		}
	}

	toStepID, err := mgr.workflowRep.GetLinearNextStepID(stepID)
	if mgr.workflowRep.IsRecordNotFoundError(err) {
		toStepID = 0
		err = nil
	} else if err != nil {
		deleteLog.WithField("err", err).Error("Failed to get next linear step of to be deleted step")
		return
	}

	if fromStep != nil {
		err = mgr.workflowRep.DeleteNextStep(fromStep.PrevID, fromStep.NextID)
		if err != nil {
			deleteLog.WithField("err", err).Error("Failed to delete nextStep pointing to to be deleted step")
			return
		}
	}

	if toStepID > 0 {
		err = mgr.workflowRep.DeleteNextStep(stepID, toStepID)
		if err != nil {
			deleteLog.WithField("err", err).Error("Failed to delete nextStep pointing from to be deleted step - ignore for now")
			err = nil
		}
	}

	err = mgr.workflowRep.DeleteStep(stepID)
	if err != nil {
		deleteLog.WithField("err", err).Error("Failed to delete step")
		return
	}

	if fromStep != nil && toStepID > 0 {
		newNextStep := &models.NextStep{
			PrevID:        fromStep.PrevID,
			NextID:        toStepID,
			DecisionLabel: fromStep.DecisionLabel,
		}
		err = mgr.workflowRep.CreateNextStep(newNextStep)
		if err != nil {
			deleteLog.WithField("err", err).Error("Failed to create new nextStep after deleting step")
			return
		}
	}

	if stepID == template.StartStep {
		if toStepID > 0 {
			template.StartStep = toStepID
		} else {
			template.StartStep = 0
		}
		err = mgr.workflowRep.UpdateTemplate(template)
		if err != nil {
			deleteLog.WithField("err", err).Error("Failed to set template startStep to 0 after deleting start step")
			return
		}
	}

	return
}

func (mgr *WorkflowManager) CreateNewRevision(oldID models.WorkflowTemplateID, revisionLabel string) (template *models.WorkflowTemplate, err error) {
	revisionLog := log.WithField("oldID", oldID)

	oldTemplate, err := mgr.GetTemplate(oldID)
	if err != nil {
		revisionLog.WithField("err", err).Error("Failed to get old template for new revision")
		return
	}

	newTemplate := &models.WorkflowTemplate{Label: revisionLabel}
	if oldTemplate.FirstRevisionID != 0 {
		newTemplate.FirstRevisionID = oldTemplate.FirstRevisionID
	} else {
		newTemplate.FirstRevisionID = oldID
	}
	revisionLog = revisionLog.WithField("newTemplate", newTemplate)

	err = mgr.CreateTemplate(newTemplate)
	if err != nil {
		revisionLog.WithField("err", err).Error("Failed to create new template for revision")
		return
	}

	if len(oldTemplate.Steps) > 0 {
		err = mgr.copySteps(newTemplate.ID, oldTemplate.Steps, 0, true, "")
		if err != nil {
			revisionLog.WithField("err", err).Error("Failed to copy steps from old to new revision template")
			mgr.workflowRep.DeleteTemplate(newTemplate.ID)
		}
	}

	return mgr.GetTemplate(newTemplate.ID)
}

func (mgr *WorkflowManager) copySteps(templateID models.WorkflowTemplateID, oldSteps []*models.Step, newStartID models.StepID, firstIsStartStep bool, decision string) (err error) {
	copyStepsLog := log.WithFields(log.Fields{"templateID": templateID, "oldSteps": oldSteps, "newStartID": newStartID, "firstIsStartStep": firstIsStartStep})

	currentPrevStep := newStartID
	for it, oldStep := range oldSteps {
		newStep := &models.Step{
			Label:  oldStep.Label,
			RoomID: oldStep.RoomID,
		}

		if it == 0 && firstIsStartStep {
			err = mgr.CreateTemplateStart(templateID, newStep)
		} else if it == 0 {
			err = mgr.AddTemplateStep(templateID, currentPrevStep, decision, newStep)
		} else {
			err = mgr.AddTemplateStep(templateID, currentPrevStep, "", newStep)
		}
		if err != nil {
			copyStepsLog.WithField("err", err).Error("Failed to create step to copy steps")
			break
		}

		for decision, steps := range oldStep.Options {
			err = mgr.copySteps(templateID, steps, newStep.ID, false, decision)
			if err != nil {
				copyStepsLog.WithField("err", err).Error("Failed to copy steps for decisions")
				continue
			}
		}

		currentPrevStep = newStep.ID
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

	err = mgr.fillExecOptions(exec)
	if err != nil {
		execLog.WithField("err", err).Error("Failed to fill exec options")
		return
	}
	return
}

func (mgr *WorkflowManager) GetExecByTrackerID(trackerID models.TrackerID) (exec *models.WorkflowExec, err error) {
	execLog := log.WithField("trackerID", trackerID)

	exec, err = mgr.workflowRep.GetRunningExecByTrackerID(trackerID)
	if err != nil {
		execLog.WithField("err", err).Error("Failed to get workflow exec by tracker")
		return
	}

	err = mgr.fillExecOptions(exec)
	if err != nil {
		execLog.WithField("err", err).Error("Failed to fill exec options")
		return
	}
	return
}

func (mgr *WorkflowManager) fillExecOptions(exec *models.WorkflowExec) (err error) {
	execOptionsLog := log.WithField("execID", exec.ID)

	infos, err := mgr.workflowRep.GetExecStepInfoForExecID(exec.ID)
	if err != nil {
		execOptionsLog.WithField("err", err).Error("Failed to get infos for exec")
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

	if tracker.Status != models.TrackerStatusIdle {
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

	timeNow := time.Now()

	exec.Status = models.ExecStatusRunning
	exec.CurrentStepID = template.StartStep
	exec.StartedOn = &timeNow
	err = mgr.workflowRep.CreateExec(exec)
	if err != nil {
		startExecLog.WithField("err", err).Warn("Failed to create workflow exec")
		return
	}

	err = GetTrackerManager().SetTrackerStatus(exec.TrackerID, models.TrackerStatusTracking)
	if err != nil {
		startExecLog.WithField("err", err).Error("Failed to set tracker to status tracking - error ignored for now")
		err = nil
	}

	startInfo := &models.ExecStepInfo{
		ExecID:    exec.ID,
		StepID:    template.StartStep,
		StartedOn: &timeNow,
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

func (mgr *WorkflowManager) ProgressToTrackerRoom(trackerID models.TrackerID, roomID models.RoomID) (err error) {
	progressLog := log.WithFields(log.Fields{"trackerID": trackerID, "roomID": roomID})

	exec, err := mgr.GetExecByTrackerID(trackerID)
	if mgr.workflowRep.IsRecordNotFoundError(err) {
		progressLog.Warn("No exec found for tracker")
		return errors.New("No exec found for tracker")
	}

	return mgr.progress(exec, nil, &roomID, progressLog)
}

func (mgr *WorkflowManager) ProgressToStep(execID models.WorkflowExecID, stepID models.StepID) (err error) {
	progressLog := log.WithFields(log.Fields{"execID": execID, "stepID": stepID})

	exec, err := mgr.GetExec(execID)
	if err != nil {
		progressLog.WithField("err", err).Error("Failed to get exec")
		return
	}

	return mgr.progress(exec, &stepID, nil, progressLog)
}

func (mgr *WorkflowManager) progress(exec *models.WorkflowExec, stepID *models.StepID, roomID *models.RoomID, progressLog *log.Entry) (err error) {
	if exec.Status != models.ExecStatusRunning {
		progressLog.Error("Workflow exec is not in status running")
		return errors.New("Workflow exec is not in status running")
	}

	template, err := mgr.GetTemplate(exec.TemplateID)
	if err != nil {
		progressLog.WithField("err", err).Error("Failed to get template of exec")
		return
	}

	updatedStepInfo := make([]*models.ExecStepInfo, 0)
	currentPassed := false
	progressToNextStep := false
	found, err := mgr.progressInSteps(exec, stepID, roomID, template.Steps, &currentPassed, &progressToNextStep, &updatedStepInfo, progressLog)
	if (!found && !*&progressToNextStep) || err != nil {
		progressLog.WithFields(log.Fields{"found": found, "err": err}).Error("Failed to progress to step")
		err = fmt.Errorf("failed to progress step: %v", err)
		return
	}

	// If the next step should be progressed to but is the last, the workflow finished
	if *&progressToNextStep {
		err = mgr.SetExecutionFinished(exec.ID)
		if err != nil {
			progressLog.WithField("err", err).Error("Failed to set execution as finished after progressing to step")
			return
		}
	}

	for _, stepInfo := range updatedStepInfo {
		err = mgr.workflowRep.UpdateExecStepInfo(stepInfo)
		if err != nil {
			progressLog.WithFields(log.Fields{"err": err, "stepInfo": stepInfo}).Error("Failed to update exec step info after progressing")
			continue
		}
	}
	return
}

func (mgr *WorkflowManager) progressInSteps(exec *models.WorkflowExec, stepID *models.StepID, roomID *models.RoomID, steps []*models.Step, currentPassed, progressToNextStep *bool, updatedStepInfo *[]*models.ExecStepInfo, progressLog *log.Entry) (bool, error) {
	timeNow := time.Now()
	for _, step := range steps {
		currentStepInfo, stepInfoOK := exec.StepInfos[step.ID]

		if step.ID == exec.CurrentStepID {
			*currentPassed = true
		}

		//Step: Found step and it is the current one => Progress to the one after
		if stepID != nil && step.ID == *stepID && exec.CurrentStepID == *stepID {
			*progressToNextStep = true
		} else if roomID != nil && *currentPassed && step.RoomID == *roomID && !*progressToNextStep { //Room: Past the current step, found step with our room => Set this step as completed, progress to the one after
			if stepInfoOK {
				currentStepInfo.CompletedOn = &timeNow
				if currentStepInfo.StartedOn == nil {
					currentStepInfo.StartedOn = &timeNow
				}
				*updatedStepInfo = append(*updatedStepInfo, currentStepInfo)
			} else {
				*updatedStepInfo = append(*updatedStepInfo, &models.ExecStepInfo{ExecID: exec.ID, StepID: step.ID, StartedOn: &timeNow, CompletedOn: &timeNow})
			}
			*progressToNextStep = true
		} else if (stepID != nil && step.ID == *stepID) || *progressToNextStep { //Step: Found step or next to progress to => Mark as started, set as current
			*progressToNextStep = false
			if stepInfoOK {
				currentStepInfo.StartedOn = &timeNow
				*updatedStepInfo = append(*updatedStepInfo, currentStepInfo)
			} else {
				*updatedStepInfo = append(*updatedStepInfo, &models.ExecStepInfo{ExecID: exec.ID, StepID: step.ID, StartedOn: &timeNow})
			}

			exec.CurrentStepID = step.ID
			err := mgr.workflowRep.UpdateExec(exec)
			if err != nil {
				progressLog.WithField("err", err).Error("Failed to set new current step id for exec")
				return true, err
			}
			return true, nil // Return that we have progressed and the next step should not be progressed to
		}

		//If the current step has passed and for room it is not the room we are searching four => Mark as skipped
		if *currentPassed && !(roomID != nil && step.RoomID == *roomID) {
			if stepInfoOK {
				currentStepInfo.Skipped = true
				*updatedStepInfo = append(*updatedStepInfo, currentStepInfo)
			} else {
				*updatedStepInfo = append(*updatedStepInfo, &models.ExecStepInfo{ExecID: exec.ID, StepID: step.ID, Skipped: true})
			}
		}

		//If we have a decision save to follow => Call recursive for these steps
		if stepInfoOK && currentStepInfo.Decision != "" {
			var found bool
			var err error
			found, err = mgr.progressInSteps(exec, stepID, roomID, step.Options[currentStepInfo.Decision], currentPassed, progressToNextStep, updatedStepInfo, progressLog)
			if err != nil {
				progressLog.WithField("err", err).Error("Failed to progress in inner steps")
				return found, err
			}
			// If we found what we are searching for and don't need to progress to next step => Exit
			if found && !*progressToNextStep {
				return true, nil
			}
		}
	}

	return false, nil
}

func (mgr *WorkflowManager) SetExecutionFinished(execID models.WorkflowExecID) (err error) {
	execFinishLog := log.WithField("execID", execID)

	exec, err := mgr.GetExec(execID)
	if err != nil {
		execFinishLog.WithField("err", err).Error("Failed to get exec")
		return
	}

	timeNow := time.Now()
	exec.CompletedOn = &timeNow
	exec.Status = models.ExecStatusCompleted
	exec.CurrentStepID = 0
	err = mgr.workflowRep.UpdateExec(exec)
	if err != nil {
		execFinishLog.WithField("err", err).Error("Failed to set exec finished")
		return
	}

	err = GetTrackerManager().SetTrackerStatus(exec.TrackerID, models.TrackerStatusIdle)
	if err != nil {
		execFinishLog.WithField("err", err).Error("Failed to set tracker to idle after finishing exec")
		return
	}

	return
}
