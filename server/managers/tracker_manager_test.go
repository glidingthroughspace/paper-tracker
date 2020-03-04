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

		trackerIdle             *models.Tracker
		trackerLearningFinished *models.Tracker

		recordNotFoundErr = errors.New("record not found")
		testErr           = errors.New("error")
	)
	const (
		sleepTimeSec                  = 5
		wrongID      models.TrackerID = 0
		id           models.TrackerID = 1
	)

	BeforeEach(func() {
		trackerManager = nil

		mockCtrl = gomock.NewController(GinkgoT())
		mockTrackerRep = mock.NewMockTrackerRepository(mockCtrl)
		mockCommandRep = mock.NewMockCommandRepository(mockCtrl)
		manager = CreateTrackerManager(mockTrackerRep, mockCommandRep, sleepTimeSec)

		trackerIdle = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusIdle}
		trackerLearningFinished = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusLearningFinished}

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
		outTracker := &models.Tracker{Label: "New Tracker", Status: models.TrackerStatusIdle}

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
		commandID := models.CommandID(1)
		outTracker := &models.Tracker{ID: id, Label: "New Tracker"}
		outCmd := &models.Command{ID: commandID, TrackerID: id, Command: models.CmdSendTrackingInformation, SleepTimeSec: 10}

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
			mockCommandRep.EXPECT().Delete(commandID).Return(nil).Times(1)
			Expect(manager.PollCommand(id)).To(Equal(outCmd))
		})

		It("PollCommand returns zero sleep time if there are commands remaining", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(outTracker, nil).Times(1)
			mockCommandRep.EXPECT().GetNextCommand(id).Return(outCmd, nil).Times(2)
			mockCommandRep.EXPECT().Delete(commandID).Return(nil).Times(1)
			Expect(manager.PollCommand(id)).To(PointTo(MatchFields(IgnoreExtras, Fields{
				"SleepTimeSec": BeEquivalentTo(0),
			})))
		})
	})

	Context("Test SetTrackerStatus", func() {
		It("SetTrackerStatus calls rep exaclty once with correct status", func() {
			mockTrackerRep.EXPECT().SetStatusByID(id, models.TrackerStatusIdle).Return(nil).Times(1)
			Expect(manager.SetTrackerStatus(id, models.TrackerStatusIdle)).To(Succeed())
		})

		It("SetTrackerStatus should return db error", func() {
			mockTrackerRep.EXPECT().SetStatusByID(id, models.TrackerStatusIdle).Return(testErr).Times(1)
			Expect(manager.SetTrackerStatus(id, models.TrackerStatusIdle)).To(MatchError(testErr))
		})
	})

	Context("Test AddTrackerCommand", func() {
		cmd := &models.Command{TrackerID: id}

		It("AddTrackerCommand calls rep exaclty once", func() {
			mockCommandRep.EXPECT().Create(cmd).Return(nil).Times(1)
			Expect(manager.AddTrackerCommand(cmd)).To(Succeed())
		})

		It("AddTrackerCommand should return db error", func() {
			mockCommandRep.EXPECT().Create(cmd).Return(testErr).Times(1)
			Expect(manager.AddTrackerCommand(cmd)).To(MatchError(testErr))
		})
	})

	Context("Test NewTrackingData", func() {
		It("NewTrackingData throws error for tracker with status LearningFinished", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearningFinished, nil).Times(1)
			Expect(manager.NewTrackingData(id, nil)).To(HaveOccurred())
		})

		It("NewTrackingData throws error for tracker with status Idle", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerIdle, nil).Times(1)
			Expect(manager.NewTrackingData(id, nil)).To(HaveOccurred())
		})
	})
})
