package managers

import (
	"errors"
	"paper-tracker/mock"
	"paper-tracker/models"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("TrackerManager", func() {
	var (
		mockTrackerRep *mock.MockTrackerRepository
		mockCommandRep *mock.MockCommandRepository
		mockCtrl       *gomock.Controller
		manager        *TrackerManager

		recordNotFoundErr = errors.New("record not found")
		testErr           = errors.New("error")
	)
	const (
		sleepTimeSec = 5
		wrongID      = 0
		id           = 1
	)

	BeforeEach(func() {
		trackerManager = nil

		mockCtrl = gomock.NewController(GinkgoT())
		mockTrackerRep = mock.NewMockTrackerRepository(mockCtrl)
		mockCommandRep = mock.NewMockCommandRepository(mockCtrl)
		manager = CreateTrackerManager(mockTrackerRep, mockCommandRep, sleepTimeSec)

		gormNotFound := func(err error) bool {
			return err == recordNotFoundErr
		}
		mockTrackerRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
		mockCommandRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()

	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Test GetAllTrackers", func() {
		outTrackers := []*models.Tracker{&models.Tracker{ID: 1, Label: "Tracker 1"}}

		It("GetAllTrackers should call get all in rep exactly once", func() {
			mockTrackerRep.EXPECT().GetAll().Return(outTrackers, nil).Times(1)
			Expect(manager.GetAllTrackers()).To(Equal(outTrackers))
		})

		It("GetAllTrackers should return db error", func() {
			mockTrackerRep.EXPECT().GetAll().Return(nil, testErr).Times(1)
			_, err := manager.GetAllTrackers()
			Expect(err).To(MatchError(testErr))
		})
	})

	Context("Test NotifyNewTracker", func() {
		outTracker := &models.Tracker{Label: "New Tracker"}

		It("NotifyNewTracker calls create in rep exactly once", func() {
			mockTrackerRep.EXPECT().Create(outTracker).Return(nil).Times(1)
			Expect(manager.NotifyNewTracker()).To(Equal(outTracker))
		})

		It("NotifyNewTracker should return db error", func() {
			mockTrackerRep.EXPECT().Create(outTracker).Return(testErr).Times(1)
			_, err := manager.NotifyNewTracker()
			Expect(err).To(MatchError(testErr))
		})
	})

	Context("Test PollCommand", func() {
		id := 1
		outTracker := &models.Tracker{ID: id, Label: "New Tracker"}
		outCmd := &models.Command{ID: 1, TrackerID: id, Command: models.CmdSendTrackingInformation, SleepTimeSec: 10}

		It("PollCommand returns correct sleep if no command in DB", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(outTracker, nil).Times(1)
			mockCommandRep.EXPECT().GetNextCommand(id).Return(nil, recordNotFoundErr).Times(1)
			Expect(manager.PollCommand(id)).To(PointTo(MatchFields(IgnoreExtras, Fields{
				"SleepTimeSec": Equal(sleepTimeSec),
				"Command":      Equal(models.CmdSleep),
			})))
		})

		It("PollCommand returns correct command from DB and deletes it", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(outTracker, nil).Times(1)
			mockCommandRep.EXPECT().GetNextCommand(id).Return(outCmd, nil).MinTimes(1)
			mockCommandRep.EXPECT().Delete(id).Return(nil).Times(1)
			Expect(manager.PollCommand(id)).To(Equal(outCmd))
		})

		It("PollCommand returns zero sleep time if there are commands remaining", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(outTracker, nil).Times(1)
			mockCommandRep.EXPECT().GetNextCommand(id).Return(outCmd, nil).Times(2)
			mockCommandRep.EXPECT().Delete(id).Return(nil).Times(1)
			Expect(manager.PollCommand(id)).To(PointTo(MatchFields(IgnoreExtras, Fields{
				"SleepTimeSec": BeEquivalentTo(0),
			})))
		})
	})
})
