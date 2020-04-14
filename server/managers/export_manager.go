package managers

import (
	"io"
	"math/rand"
	"paper-tracker/models"
	"paper-tracker/repositories"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

var exportManager *ExportManager

type ExportManager struct {
	workflowTemplateRep repositories.WorkflowTemplateRepository
	workflowExecRep     repositories.WorkflowExecRepository
}

func CreateExportManager(workflowTemplateRep repositories.WorkflowTemplateRepository, workflowExecRep repositories.WorkflowExecRepository) *ExportManager {
	if trackingManager != nil {
		return exportManager
	}

	rand.Seed(time.Now().UnixNano())

	exportManager = &ExportManager{
		workflowTemplateRep: workflowTemplateRep,
		workflowExecRep:     workflowExecRep,
	}

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
	templates, err := mgr.workflowTemplateRep.GetAllTemplates()
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

	execs, err := mgr.workflowExecRep.GetExecsByTemplateID(template.ID)
	if err != nil {
		log.WithError(err).WithField("templateID", template.ID).Error("Failed to get all execs for export")
		return
	}
	for it, exec := range execs {
		sheet.Cell(it+1, 0).SetString(exec.Label)
		sheet.Cell(it+1, 1).SetString(exec.Status.String())
		sheet.Cell(it+1, 2).SetDateTime(*exec.StartedOn)
		sheet.Cell(it+1, 3).SetDateTime(*exec.CompletedOn)
	}

	return
}
