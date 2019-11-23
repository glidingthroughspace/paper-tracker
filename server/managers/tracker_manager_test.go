package managers

import (
	"errors"
	"paper-tracker/mock"
	"paper-tracker/models"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	log "github.com/sirupsen/logrus"
)

var _ = Describe("TrackerManager", func() {
	var (
		mockTrackerRep    *mock.MockTrackerRepository
		mockCommandRep    *mock.MockCommandRepository
		mockScanResultRep *mock.MockScanResultRepository
		mockCtrl          *gomock.Controller
		manager           *TrackerManager
	)
	const sleepTimeSec = 5
	const sleepBetweenLearnSec = 1
	const learnCount = 2

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockTrackerRep = mock.NewMockTrackerRepository(mockCtrl)
		mockCommandRep = mock.NewMockCommandRepository(mockCtrl)
		mockScanResultRep = mock.NewMockScanResultRepository(mockCtrl)
		manager = CreateTrackerManager(mockTrackerRep, mockCommandRep, mockScanResultRep, sleepTimeSec, learnCount, sleepBetweenLearnSec)

		gormNotFound := func(err error) bool {
			return gorm.IsRecordNotFoundError(err)
		}
		mockTrackerRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
		mockCommandRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
		mockScanResultRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Test GetAllTrackers", func() {
		outTrackers := []*models.Tracker{&models.Tracker{ID: 1, Label: "Tracker 1"}}
		expErr := errors.New("error")

		It("GetAllTrackers should call get all in rep exactly once", func() {
			mockTrackerRep.EXPECT().GetAll().Return(outTrackers, nil).Times(1)
			Expect(manager.GetAllTrackers()).To(Equal(outTrackers))
		})

		It("GetAllTrackers should return db error", func() {
			mockTrackerRep.EXPECT().GetAll().Return(nil, expErr).Times(1)
			_, err := manager.GetAllTrackers()
			Expect(err).To(MatchError(expErr))
		})
	})

	Context("Test NotifyNewTracker", func() {
		outTracker := &models.Tracker{Label: "New Tracker"}
		expErr := errors.New("error")

		It("NotifyNewTracker calls create in rep exactly once", func() {
			mockTrackerRep.EXPECT().Create(outTracker).Return(nil).Times(1)
			Expect(manager.NotifyNewTracker()).To(Equal(outTracker))
		})

		It("NotifyNewTracker should return db error", func() {
			mockTrackerRep.EXPECT().Create(outTracker).Return(expErr).Times(1)
			_, err := manager.NotifyNewTracker()
			Expect(err).To(MatchError(expErr))
		})
	})

	Context("Test PollCommand", func() {
		id := 1
		outTracker := &models.Tracker{ID: id, Label: "New Tracker"}
		outCmd := &models.Command{ID: 1, TrackerID: id, Command: models.CmdSendTrackingInformation, SleepTimeSec: 10}

		It("PollCommand returns correct sleep if no command in DB", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(outTracker, nil).Times(1)
			mockCommandRep.EXPECT().GetNextCommand(id).Return(nil, gorm.ErrRecordNotFound).Times(1)
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

	Context("Test StartLearning", func() {
		wrongID := 0
		id := 1
		var trackerIdle *models.Tracker
		var trackerLearning *models.Tracker
		var cmdCreateCall *gomock.Call
		var trackerUpdateCall *gomock.Call
		trackCmd := &models.Command{TrackerID: id, Command: models.CmdSendTrackingInformation, SleepTimeSec: sleepBetweenLearnSec}
		testLogger := log.WithField("unit_test", true)

		BeforeEach(func() {
			trackerIdle = &models.Tracker{ID: id, Label: "New Tracker", Status: models.StatusIdle}
			trackerLearning = &models.Tracker{ID: id, Label: "New Tracker", Status: models.StatusLearning}
			cmdCreateCall = mockCommandRep.EXPECT().Create(trackCmd).Return(nil).AnyTimes()
			trackerUpdateCall = mockTrackerRep.EXPECT().Update(gomock.Any()).AnyTimes()
		})

		It("StartLearning returns error if tracker does not exist", func() {
			mockTrackerRep.EXPECT().GetByID(wrongID).Return(nil, gorm.ErrRecordNotFound).Times(1)
			_, err := manager.StartLearning(wrongID)
			Expect(err).To(MatchError(gorm.ErrRecordNotFound))
		})

		It("StartLearning returns error if tracker is not in idle mode", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearning, nil).Times(1)
			_, err := manager.StartLearning(id)
			Expect(err).To(HaveOccurred())
		})

		It("StartLearning return correct total learn time", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerIdle, nil).Times(1)
			Expect(manager.StartLearning(id)).To(Equal(learnCount * sleepBetweenLearnSec))
		})

		It("StartLearning inserts first command", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerIdle, nil).Times(1)
			cmdCreateCall.MinTimes(1)
			_, err := manager.StartLearning(id)
			Expect(err).To(Succeed())
			time.Sleep(10 * time.Millisecond)
		})

		Context("Test learningRoutine", func() {
			It("learningRoutine sets tracker status to learning and back to idle", func() {
				trackerUpdateCall.Times(2)
				manager.learningRoutine(trackerIdle, testLogger)
			})
		})

		Context("Test learningSendTrackingCmds", func() {
			It("learningSendTrackingCmds insert correct amounts of tracking commands", func() {
				cmdCreateCall.Times(learnCount)
				manager.learningCreateTrackingCmds(trackerLearning, testLogger)
			})
		})
	})

	Context("Test NewTrackingData", func() {
		//TODO
	})
})
