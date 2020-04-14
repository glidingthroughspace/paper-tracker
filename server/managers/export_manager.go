package managers

import (
	"io"
	"math/rand"
	"paper-tracker/models"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

var exportManager *ExportManager

type ExportManager struct {
}

func CreateExportManager() *ExportManager {
	if trackingManager != nil {
		return exportManager
	}

	rand.Seed(time.Now().UnixNano())

	exportManager = &ExportManager{}

	return exportManager
}

func GetExportManager() *ExportManager {
	return exportManager
}

func (mgr *ExportManager) GenerateExport(writer io.Writer) (err error) {
	file := xlsx.NewFile()

	err = mgr.fillExportFile(file)
	if err != nil {
		log.WithError(err).Error("Failed to generate export")
		return
	}

	err = file.Write(writer)
	return
}

type templateExport struct {
	template                 *models.WorkflowTemplate
	numExecutions            int
	percentageCompleted      int
	meanCompletedExecTimeHrs float64
}

func (mgr *ExportManager) fillExportFile(file *xlsx.File) (err error) {
	templates, err := GetWorkflowTemplateManager().GetAllTemplates()
	if err != nil {
		log.WithError(err).Error("Failed to get templates to export")
		return
	}

	// List to gather templates with export info
	tmplExports := make(map[models.WorkflowTemplateID]*templateExport, len(templates))

	// Go through all templates and fill their own template and save needed data for list
	for _, template := range templates {
		sheet, err := file.AddSheet(template.Label)
		if err != nil {
			log.WithError(err).WithField("templateID", template.ID).Error("Failed to create template sheet for export")
			continue
		}

		err, tmplExport := mgr.fillExportSheet(template, sheet)
		if err != nil {
			log.WithError(err).WithField("templateID", template.ID).Error("Failed to fill template sheet for export")
			continue
		}
		tmplExports[template.ID] = tmplExport
	}

	// Assemble revision tab based on list

	return
}

func (mgr *ExportManager) fillExportSheet(template *models.WorkflowTemplate, sheet *xlsx.Sheet) (err error, export *templateExport) {
	sheet.Cell(0, 0).SetString("Label of execution")
	sheet.Cell(0, 1).SetString("Status")
	sheet.Cell(0, 2).SetString("Start Time")
	sheet.Cell(0, 3).SetString("End Time")

	execs, err := GetWorkflowExecManager().GetExecsByTemplate(template.ID)
	if err != nil {
		log.WithError(err).WithField("templateID", template.ID).Error("Failed to get all execs for export")
		return
	}

	stepInfoCols := make(map[models.StepID]int)
	lastStepEndCol := 4
	for it, exec := range execs {
		row := it + 1
		sheet.Cell(row, 0).SetString(exec.Label)
		sheet.Cell(row, 1).SetString(exec.Status.String())
		if exec.StartedOn != nil {
			sheet.Cell(row, 2).SetDateTime(*exec.StartedOn)
		}
		if exec.CompletedOn != nil {
			sheet.Cell(row, 3).SetDateTime(*exec.CompletedOn)
		}

		for _, stepInfo := range exec.StepInfos {
			col := -1
			if savedCol, ok := stepInfoCols[stepInfo.StepID]; ok {
				col = savedCol
			} else {
				col = lastStepEndCol + 1
				lastStepEndCol = col + 5
				stepInfoCols[stepInfo.StepID] = col
				mgr.fillStepInfoHeader(sheet, col, template.ID, stepInfo.StepID)
			}
			sheet.Cell(row, col).SetBool(true)
			if stepInfo.StartedOn != nil {
				sheet.Cell(row, col+1).SetDateTime(*stepInfo.StartedOn)
			}
			if stepInfo.CompletedOn != nil {
				sheet.Cell(row, col+2).SetDateTime(*stepInfo.CompletedOn)
			}
			sheet.Cell(row, col+3).SetString(stepInfo.Decision)
			sheet.Cell(row, col+4).SetBool(stepInfo.Skipped)
		}
	}

	return
}

func (mgr *ExportManager) fillStepInfoHeader(sheet *xlsx.Sheet, col int, templateID models.WorkflowTemplateID, stepID models.StepID) {
	step, err := GetWorkflowTemplateManager().GetStepByID(templateID, stepID)
	if err != nil {
		log.WithError(err).WithField("stepID", stepID).Error("Failed to get step to fill stepInfoHeader for export")
		return
	}
	roomLabel := ""
	if room, err := GetRoomManager().GetRoomByID(step.RoomID); err == nil {
		roomLabel = room.Label
	} else {
		log.WithError(err).WithField("roomID", stepID).Error("Failed to get room to fill stepInfoHeader for export - ignore for now")
	}

	sheet.Cell(0, col).SetString(step.Label + " (" + roomLabel + ") Part of Exec")
	sheet.Cell(0, col+1).SetString(step.Label + "Start Time")
	sheet.Cell(0, col+2).SetString(step.Label + " End Time")
	sheet.Cell(0, col+3).SetString(step.Label + " Decision")
	sheet.Cell(0, col+4).SetString(step.Label + " Skipped")
}
