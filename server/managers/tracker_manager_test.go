package managers_test

import (
	. "paper-tracker/managers"
	"paper-tracker/mock"
	"paper-tracker/models"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TrackerManager", func() {
	var (
		mockTrackerRep *mock.MockTrackerRepository
		mockCommandRep *mock.MockCommandRepository
		mockCtrl       *gomock.Controller
		manager        *TrackerManager
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockTrackerRep = mock.NewMockTrackerRepository(mockCtrl)
		mockCommandRep = mock.NewMockCommandRepository(mockCtrl)
		manager = CreateTrackerManager(mockTrackerRep, mockCommandRep, 5)

	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Test NotifyNewTracker", func() {
		var createCall *gomock.Call
		inTracker := &models.Tracker{Label: "New Tracker"}
		BeforeEach(func() {
			createCall = mockTrackerRep.EXPECT().Create(inTracker).Return(nil).AnyTimes()
		})
		It("NotifyNewTracker calls create in rep exactly once", func() {
			Expect(manager.NotifyNewTracker()).Should(Equal(inTracker))
			createCall.Times(1)
		})
	})
})
