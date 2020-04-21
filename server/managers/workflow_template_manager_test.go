package managers

import (
	"errors"
	"paper-tracker/mock"
	"paper-tracker/models"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("WorkflowTemplateManager", func() {
	var (
		mockWorkflowTemplateRep *mock.MockWorkflowTemplateRepository
		mockCtrl                *gomock.Controller
		manager                 *WorkflowTemplateManagerImpl
		mockWorkflowExecMgr     *mock.MockWorkflowExecManager

		recordNotFoundErr = errors.New("record not found")
		testErr           = errors.New("error")
	)
	const (
		sleepTimeSec = 5
		wrongID      = 0
		id           = 1
	)

	BeforeEach(func() {
		workflowTemplateManager = nil
		workflowExecManager = nil

		mockCtrl = gomock.NewController(GinkgoT())
		mockWorkflowTemplateRep = mock.NewMockWorkflowTemplateRepository(mockCtrl)
		manager = CreateWorkflowTemplateManager(mockWorkflowTemplateRep).(*WorkflowTemplateManagerImpl)

		mockWorkflowExecMgr = mock.NewMockWorkflowExecManager(mockCtrl)
		workflowExecManager = mockWorkflowExecMgr

		gormNotFound := func(err error) bool {
			return err == recordNotFoundErr
		}
		mockWorkflowTemplateRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Test CreateWorkflowTemplate", func() {
		workflow := &models.WorkflowTemplate{ID: 0, Label: "TestWorkflow"}

		It("CreateTemplate calls repo create exactly once", func() {
			mockWorkflowTemplateRep.EXPECT().CreateTemplate(workflow).Return(nil).Times(1)
			Expect(manager.CreateTemplate(workflow)).To(Succeed())
		})

		It("CreateTemplate return db error", func() {
			mockWorkflowTemplateRep.EXPECT().CreateTemplate(workflow).Return(testErr).AnyTimes()
			Expect(manager.CreateTemplate(workflow)).To(MatchError(testErr))
		})

		It("CreateTemplate sets ID to 0", func() {
			workflow.ID = 1
			mockWorkflowTemplateRep.EXPECT().CreateTemplate(workflow).Return(nil).AnyTimes()
			manager.CreateTemplate(workflow)
			Expect(workflow.ID).To(BeEquivalentTo(0))
		})
	})

	Context("Test CreateTemplateStart", func() {
		workflow := &models.WorkflowTemplate{ID: 1, Label: "TestWorkflow"}
		step := &models.Step{ID: 1, Label: "TestStep"}

		It("CreateTemplateStart calls repo create exactly once, gets workflow and inserts startstep", func() {
			mockWorkflowTemplateRep.EXPECT().CreateStep(step).Return(nil).Times(1)
			mockWorkflowTemplateRep.EXPECT().GetTemplateByID(workflow.ID).Return(workflow, nil).Times(1)
			mockWorkflowExecMgr.EXPECT().GetExecCountByTemplate(workflow.ID).Return(0, nil).Times(1)
			mockWorkflowTemplateRep.EXPECT().UpdateTemplate(workflow).Return(nil).Times(1)

			Expect(manager.CreateTemplateStart(workflow.ID, step)).To(Succeed())
		})

		//TODO: Add more tests
	})
})
