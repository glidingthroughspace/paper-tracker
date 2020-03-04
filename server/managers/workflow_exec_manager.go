package managers

import (
	"errors"
	"fmt"
	"paper-tracker/models"
	"paper-tracker/repositories"
	"time"

	log "github.com/sirupsen/logrus"
)

var workflowExecManager *WorkflowExecManager

type WorkflowExecManager struct {
	workflowRep repositories.WorkflowRepository
}

func CreateWorkflowExecManager(workflowRep repositories.WorkflowRepository) *WorkflowExecManager {
	if workflowExecManager != nil {
		return workflowExecManager
	}

	workflowExecManager = &WorkflowExecManager{
		workflowRep: workflowRep,
	}

	return workflowExecManager
}

func GetWorkflowExecManager() *WorkflowExecManager {
	return workflowExecManager
}

func (mgr *WorkflowExecManager) GetAllExec() (execs []*models.WorkflowExec, err error) {
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

func (mgr *WorkflowExecManager) GetExec(execID models.WorkflowExecID) (exec *models.WorkflowExec, err error) {
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

func (mgr *WorkflowExecManager) GetExecByTrackerID(trackerID models.TrackerID) (exec *models.WorkflowExec, err error) {
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

func (mgr *WorkflowExecManager) fillExecOptions(exec *models.WorkflowExec) (err error) {
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

func (mgr *WorkflowExecManager) StartExecution(exec *models.WorkflowExec) (err error) {
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

	template, err := GetWorkflowTemplateManager().GetTemplate(exec.TemplateID)
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

func (mgr *WorkflowExecManager) ProgressToTrackerRoom(trackerID models.TrackerID, roomID models.RoomID) (err error) {
	progressLog := log.WithFields(log.Fields{"trackerID": trackerID, "roomID": roomID})

	exec, err := mgr.GetExecByTrackerID(trackerID)
	if mgr.workflowRep.IsRecordNotFoundError(err) {
		progressLog.Warn("No exec found for tracker")
		return errors.New("No exec found for tracker")
	}

	return mgr.progress(exec, nil, &roomID, progressLog)
}

func (mgr *WorkflowExecManager) ProgressToStep(execID models.WorkflowExecID, stepID models.StepID) (err error) {
	progressLog := log.WithFields(log.Fields{"execID": execID, "stepID": stepID})

	exec, err := mgr.GetExec(execID)
	if err != nil {
		progressLog.WithField("err", err).Error("Failed to get exec")
		return
	}

	return mgr.progress(exec, &stepID, nil, progressLog)
}

func (mgr *WorkflowExecManager) progress(exec *models.WorkflowExec, stepID *models.StepID, roomID *models.RoomID, progressLog *log.Entry) (err error) {
	if exec.Status != models.ExecStatusRunning {
		progressLog.Error("Workflow exec is not in status running")
		return errors.New("Workflow exec is not in status running")
	}

	template, err := GetWorkflowTemplateManager().GetTemplate(exec.TemplateID)
	if err != nil {
		progressLog.WithField("err", err).Error("Failed to get template of exec")
		return
	}

	updatedStepInfo := make([]*models.ExecStepInfo, 0)
	currentPassed := false
	progressToNextStep := false
	found, err := mgr.progressInSteps(exec, stepID, roomID, template.Steps, &currentPassed, &progressToNextStep, &updatedStepInfo, progressLog)
	if (!found && !progressToNextStep) || err != nil {
		progressLog.WithFields(log.Fields{"found": found, "err": err}).Error("Failed to progress to step")
		err = fmt.Errorf("failed to progress step: %v", err)
		return
	}

	// If the next step should be progressed to but is the last, the workflow finished
	if progressToNextStep {
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

func (mgr *WorkflowExecManager) progressInSteps(exec *models.WorkflowExec, stepID *models.StepID, roomID *models.RoomID, steps []*models.Step, currentPassed, progressToNextStep *bool, updatedStepInfo *[]*models.ExecStepInfo, progressLog *log.Entry) (bool, error) {
	timeNow := time.Now()
	for _, step := range steps {
		currentStepInfo := exec.StepInfos[step.ID]

		if step.ID == exec.CurrentStepID {
			*currentPassed = true
		}

		//Searching for step: Found the step we are searching for and it also is the current one => Progress to the one after
		if stepID != nil && step.ID == *stepID && exec.CurrentStepID == *stepID {
			*progressToNextStep = true
		} else if roomID != nil && *currentPassed && step.RoomID == *roomID && !*progressToNextStep { //Serching for room: We are past the current step of the exec, found the step with the roomID we are searching for => Set this step as completed, progress to the one after
			mgr.markStepAsCompleted(exec, step, currentStepInfo, timeNow, updatedStepInfo)
			*progressToNextStep = true
		} else if (stepID != nil && step.ID == *stepID) || *progressToNextStep { //Searching for step: Found step with stepID we are searching for or we should progress to the next step => Mark as started, set as current step of the exec
			*progressToNextStep = false
			err := mgr.markStepAsCurrent(exec, step, timeNow, currentStepInfo, updatedStepInfo, progressLog)
			return true, err
		}

		//If the current step of the exec has passed and in case we are searching for a room the step has not the room we are searching for => Mark this step as skipped
		if *currentPassed && !(roomID != nil && step.RoomID == *roomID) {
			mgr.markStepAsSkipped(exec, step, timeNow, currentStepInfo, updatedStepInfo)
		}

		//If we have a decision saved to follow for this step => Call recursive for these steps
		if currentStepInfo != nil && currentStepInfo.Decision != "" {
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

func (mgr *WorkflowExecManager) markStepAsCompleted(exec *models.WorkflowExec, step *models.Step, currentStepInfo *models.ExecStepInfo, timeNow time.Time, updatedStepInfo *[]*models.ExecStepInfo) {
	if currentStepInfo != nil {
		currentStepInfo.CompletedOn = &timeNow
		if currentStepInfo.StartedOn == nil {
			currentStepInfo.StartedOn = &timeNow
		}
		*updatedStepInfo = append(*updatedStepInfo, currentStepInfo)
	} else {
		*updatedStepInfo = append(*updatedStepInfo, &models.ExecStepInfo{ExecID: exec.ID, StepID: step.ID, StartedOn: &timeNow, CompletedOn: &timeNow})
	}
}

func (mgr *WorkflowExecManager) markStepAsCurrent(exec *models.WorkflowExec, step *models.Step, timeNow time.Time, currentStepInfo *models.ExecStepInfo, updatedStepInfo *[]*models.ExecStepInfo, progressLog *log.Entry) (err error) {
	if currentStepInfo != nil {
		currentStepInfo.StartedOn = &timeNow
		*updatedStepInfo = append(*updatedStepInfo, currentStepInfo)
	} else {
		*updatedStepInfo = append(*updatedStepInfo, &models.ExecStepInfo{ExecID: exec.ID, StepID: step.ID, StartedOn: &timeNow})
	}

	exec.CurrentStepID = step.ID
	err = mgr.workflowRep.UpdateExec(exec)
	if err != nil {
		progressLog.WithField("err", err).Error("Failed to set new current step id for exec")
		return
	}
	return
}

func (mgr *WorkflowExecManager) markStepAsSkipped(exec *models.WorkflowExec, step *models.Step, timeNow time.Time, currentStepInfo *models.ExecStepInfo, updatedStepInfo *[]*models.ExecStepInfo) {
	if currentStepInfo != nil {
		currentStepInfo.Skipped = true
		*updatedStepInfo = append(*updatedStepInfo, currentStepInfo)
	} else {
		*updatedStepInfo = append(*updatedStepInfo, &models.ExecStepInfo{ExecID: exec.ID, StepID: step.ID, Skipped: true})
	}
}

func (mgr *WorkflowExecManager) SetExecutionFinished(execID models.WorkflowExecID) (err error) {
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

func (mgr *WorkflowExecManager) CancelExec(execID models.WorkflowExecID) (err error) {
	cancelLog := log.WithField("execID", execID)

	exec, err := mgr.GetExec(execID)
	if err != nil {
		cancelLog.WithField("err", err).Error("Could not get exec to cancel")
		return
	}

	if exec.Status != models.ExecStatusRunning {
		cancelLog.Error("Exec is not in running status")
		return errors.New("Exec is not in running status")
	}

	exec.Status = models.ExecStatusCanceled
	err = mgr.workflowRep.UpdateExec(exec)
	if err != nil {
		cancelLog.WithField("err", err).Error("Failed to update exec")
		return
	}

	err = GetTrackerManager().SetTrackerStatus(exec.TrackerID, models.TrackerStatusIdle)
	if err != nil {
		cancelLog.WithFields(log.Fields{"err": err, "trackerID": exec.TrackerID}).Error("Failed to set tracker of exec to idle")
		return
	}
	return
}