package managers

import (
	"errors"
	"paper-tracker/mock"
	"paper-tracker/models"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TrackerManager", func() {
	var (
		mockWorkflowRep *mock.MockWorkflowRepository
		mockCtrl        *gomock.Controller
		manager         *WorkflowManager

		recordNotFoundErr = errors.New("record not found")
		testErr           = errors.New("error")
	)
	const (
		sleepTimeSec = 5
		wrongID      = 0
		id           = 1
	)

	BeforeEach(func() {
		workflowManager = nil

		mockCtrl = gomock.NewController(GinkgoT())
		mockWorkflowRep = mock.NewMockWorkflowRepository(mockCtrl)
		manager = CreateWorkflowManager(mockWorkflowRep)

		gormNotFound := func(err error) bool {
			return err == recordNotFoundErr
		}
		mockWorkflowRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Test CreateWorkflow", func() {
		workflow := &models.WorkflowTemplate{ID: 0, Label: "TestWorkflow"}

		It("CreateWorkflow calls repo create exactly once", func() {
			mockWorkflowRep.EXPECT().CreateTemplate(workflow).Return(nil).Times(1)
			Expect(manager.CreateTemplate(workflow)).To(Succeed())
		})

		It("CreateWorkflow return db error", func() {
			mockWorkflowRep.EXPECT().CreateTemplate(workflow).Return(testErr).AnyTimes()
			Expect(manager.CreateTemplate(workflow)).To(MatchError(testErr))
		})

		It("CreateWorkflow sets ID to 0", func() {
			workflow.ID = 1
			mockWorkflowRep.EXPECT().CreateTemplate(workflow).Return(nil).AnyTimes()
			manager.CreateTemplate(workflow)
			Expect(workflow.ID).To(BeEquivalentTo(0))
		})
	})

	Context("Test CreateWorkflowStart", func() {
		workflow := &models.WorkflowTemplate{ID: 1, Label: "TestWorkflow"}
		step := &models.Step{ID: 1, Label: "TestStep"}

		It("CreateWorkflowStart calls repo create exactly once, gets workflow and inserts startstep", func() {
			mockWorkflowRep.EXPECT().CreateStep(step).Return(nil).Times(1)
			mockWorkflowRep.EXPECT().GetTemplateByID(workflow.ID).Return(workflow, nil).Times(1)
			mockWorkflowRep.EXPECT().UpdateTemplate(workflow).Return(nil).Times(1)

			Expect(manager.CreateTemplateStart(workflow.ID, step)).To(Succeed())
		})

		//TODO: Add more tests
	})
})
