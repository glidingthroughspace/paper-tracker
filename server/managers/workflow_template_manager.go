package managers

import (
	"errors"
	"paper-tracker/models"
	"paper-tracker/repositories"

	log "github.com/sirupsen/logrus"
)

var workflowTemplateManager *WorkflowTemplateManager

type WorkflowTemplateManager struct {
	workflowRep repositories.WorkflowTemplateRepository
}

func CreateWorkflowTemplateManager(workflowRep repositories.WorkflowTemplateRepository) *WorkflowTemplateManager {
	if workflowTemplateManager != nil {
		return workflowTemplateManager
	}

	workflowTemplateManager = &WorkflowTemplateManager{
		workflowRep: workflowRep,
	}

	return workflowTemplateManager
}

func GetWorkflowTemplateManager() *WorkflowTemplateManager {
	return workflowTemplateManager
}

func (mgr *WorkflowTemplateManager) CreateTemplate(template *models.WorkflowTemplate) (err error) {
	template.ID = 0
	err = mgr.workflowRep.CreateTemplate(template)
	if err != nil {
		log.WithFields(log.Fields{"template": template, "err": err}).Error("Failed to create template")
		return
	}
	return
}

func (mgr *WorkflowTemplateManager) CreateTemplateStart(templateID models.WorkflowTemplateID, step *models.Step) (err error) {
	workflowStartLog := log.WithFields(log.Fields{"templateID": templateID, "step": step})

	template, err := mgr.GetTemplate(templateID)
	if err != nil || template.StepEditingLocked {
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
func (mgr *WorkflowTemplateManager) AddTemplateStep(templateID models.WorkflowTemplateID, prevStepID models.StepID, decisionLabel string, step *models.Step) (err error) {
	addStepLog := log.WithFields(log.Fields{"prevStepID": prevStepID, "step": step})

	template, err := mgr.GetTemplate(templateID)
	if err != nil || template.StepEditingLocked {
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

func (mgr *WorkflowTemplateManager) GetAllTemplates() (templates []*models.WorkflowTemplate, err error) {
	templates, err = mgr.workflowRep.GetAllTemplates()
	if err != nil {
		log.WithField("err", err).Error("Failed to get all raw workflows")
		return
	}

	for _, template := range templates {
		err = mgr.fillTemplateInfo(template)
		if err != nil {
			log.WithFields(log.Fields{"err": err, "templateID": template.ID}).Error("Failed to fill workflow infos for list")
			continue
		}
	}
	return
}

func (mgr *WorkflowTemplateManager) GetTemplate(templateID models.WorkflowTemplateID) (template *models.WorkflowTemplate, err error) {
	getWorkflowLog := log.WithField("templateID", templateID)

	template, err = mgr.workflowRep.GetTemplateByID(templateID)
	if mgr.workflowRep.IsRecordNotFoundError(err) {
		getWorkflowLog.Warn("Template not found")
		return
	} else if err != nil {
		getWorkflowLog.WithField("err", err).Error("Failed to get template")
		return
	}

	err = mgr.fillTemplateInfo(template)
	return
}

func (mgr *WorkflowTemplateManager) fillTemplateInfo(template *models.WorkflowTemplate) (err error) {
	infoLog := log.WithField("templateID", template.ID)

	execCount, err := GetWorkflowExecManager().GetExecCountByTemplate(template.ID)
	if err != nil {
		infoLog.WithField("err", err).Error("Failed to get execs of template - ignore for now")
		err = nil
	}
	if execCount > 0 {
		template.StepEditingLocked = true
	} else {
		template.StepEditingLocked = false
	}

	template.Steps, err = mgr.getStepsFromStart(template.ID, template.StartStep, infoLog)
	if err != nil {
		infoLog.WithField("err", err).Error("Failed to get steps")
		return
	}
	return
}

func (mgr *WorkflowTemplateManager) getStepsFromStart(templateID models.WorkflowTemplateID, startStepID models.StepID, getLog *log.Entry) (steps []*models.Step, err error) {
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

func (mgr *WorkflowTemplateManager) GetStepByID(templateID models.WorkflowTemplateID, stepID models.StepID) (step *models.Step, err error) {
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

func (mgr *WorkflowTemplateManager) UpdateStep(templateID models.WorkflowTemplateID, step *models.Step) (err error) {
	updateLog := log.WithFields(log.Fields{"templateID": templateID, "step": step})

	template, err := mgr.GetTemplate(templateID)
	if err != nil || template.StepEditingLocked {
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

func (mgr *WorkflowTemplateManager) DeleteStep(templateID models.WorkflowTemplateID, stepID models.StepID) (err error) {
	deleteLog := log.WithFields(log.Fields{"templateID": templateID, "stepID": stepID})

	template, err := mgr.GetTemplate(templateID)
	if err != nil || template.StepEditingLocked {
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

func (mgr *WorkflowTemplateManager) CreateNewRevision(oldID models.WorkflowTemplateID, revisionLabel string) (template *models.WorkflowTemplate, err error) {
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

func (mgr *WorkflowTemplateManager) copySteps(templateID models.WorkflowTemplateID, oldSteps []*models.Step, newStartID models.StepID, firstIsStartStep bool, decision string) (err error) {
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

func (mgr *WorkflowTemplateManager) NumberOfStepsReferringToRoom(roomID models.RoomID) (int, error) {
	steps, err := mgr.workflowRep.GetStepsByRoomID(roomID)
	if err != nil {
		log.WithFields(log.Fields{"roomID": roomID, "err": err}).Error("Failed to get all steps by room id")
		return -1, err
	}
	return len(steps), nil
}
