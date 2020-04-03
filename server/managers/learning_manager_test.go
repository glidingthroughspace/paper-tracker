package managers

import (
	"errors"
	"paper-tracker/mock"
	"paper-tracker/models"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

var _ = Describe("LearningManager", func() {
	var (
		mockTrackerRep    *mock.MockTrackerRepository
		mockScanResultRep *mock.MockScanResultRepository
		mockRoomRep       *mock.MockRoomRepository
		mockTemplateRep   *mock.MockWorkflowTemplateRepository
		mockCtrl          *gomock.Controller
		manager           *LearningManager

		trackerIdle             *models.Tracker
		trackerLearning         *models.Tracker
		trackerLearningFinished *models.Tracker

		scanRes       []*models.ScanResult
		scanResWithID []*models.ScanResult
		scanResSSIDs  []string

		recordNotFoundErr = errors.New("record not found")
		testErr           = errors.New("error")
	)
	const (
		sleepBetweenLearnSec                  = 1
		learnCount                            = 1
		wrongID              models.TrackerID = 0
		wrongRoomID          models.RoomID    = 0
		id                   models.TrackerID = 1
		roomID               models.RoomID    = 1
	)

	BeforeEach(func() {
		learningManager = nil
		// TODO: Mock managers
		trackerManager = nil
		roomManager = nil
		workflowTemplateManager = nil

		mockCtrl = gomock.NewController(GinkgoT())
		mockScanResultRep = mock.NewMockScanResultRepository(mockCtrl)
		mockTrackerRep = mock.NewMockTrackerRepository(mockCtrl)
		mockRoomRep = mock.NewMockRoomRepository(mockCtrl)
		mockTemplateRep = mock.NewMockWorkflowTemplateRepository(mockCtrl)
		manager = CreateLearningManager(mockScanResultRep, learnCount, sleepBetweenLearnSec)
		CreateTrackerManager(mockTrackerRep, 5, 5, 5, 5, 5)
		CreateRoomManager(mockRoomRep)
		CreateWorkflowTemplateManager(mockTemplateRep)

		gormNotFound := func(err error) bool {
			return gorm.IsRecordNotFoundError(err)
		}
		mockScanResultRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
		mockTrackerRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
		mockRoomRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()

		trackerIdle = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusIdle}
		trackerLearning = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusLearning}
		trackerLearningFinished = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusLearningFinished}

		scanRes = []*models.ScanResult{
			&models.ScanResult{SSID: "Test0", BSSID: "aa:bb:cc:dd:ee:ff", RSSI: -50},
			&models.ScanResult{SSID: "Test1", BSSID: "aa:bb:cc:dd:ee:ff", RSSI: -40},
		}
		scanResWithID = []*models.ScanResult{
			&models.ScanResult{TrackerID: id, SSID: "Test0", BSSID: "aa:bb:cc:dd:ee:ff", RSSI: -50},
			&models.ScanResult{TrackerID: id, SSID: "Test1", BSSID: "aa:bb:cc:dd:ee:ff", RSSI: -40},
		}
		scanResSSIDs = []string{"Test0", "Test1"}
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Test StartLearning", func() {
		var trackerSetStatusCall *gomock.Call
		testLogger := log.WithField("unit_test", true)

		BeforeEach(func() {
			trackerSetStatusCall = mockTrackerRep.EXPECT().SetStatusByID(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
		})

		// Somehow need to fix testing with goroutines...
		/*It("StartLearning returns error if tracker does not exist", func() {
			mockTrackerRep.EXPECT().GetByID(wrongID).Return(nil, recordNotFoundErr).Times(1)
			_, err := manager.StartLearning(wrongID)
			Expect(err).To(MatchError(recordNotFoundErr))
		})

		It("StartLearning returns error if tracker is not in idle mode", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearning, nil).Times(1)
			_, err := manager.StartLearning(id)
			Expect(err).To(HaveOccurred())
		})

		It("StartLearning return correct total learn time", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerIdle, nil).Times(1)
			mockScanResultRep.EXPECT().DeleteForTracker(id).Return(nil).Times(1)
			Expect(manager.StartLearning(id)).To(Equal(1)) // int(1*1*1.2)
		})*/

		Context("Test learningRoutine", func() {
			It("learningRoutine sets tracker status to learning and learning finished", func() {
				trackerSetStatusCall.Times(2)
				mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearning, nil).AnyTimes()
				manager.learningRoutine(id, testLogger)
			})
		})
	})

	Context("Test newLearningTrackingData", func() {
		It("newLearningTrackingData creates all scan results in db", func() {
			mockScanResultRep.EXPECT().CreateAll(gomock.Any()).Return(nil).Times(1)
			Expect(manager.newLearningTrackingData(id, scanRes)).To(Succeed())
		})

		It("newLearningTrackingData add proper trackerID to scan results", func() {
			mockScanResultRep.EXPECT().CreateAll(scanResWithID).Return(nil).Times(1)
			Expect(manager.newLearningTrackingData(id, scanRes)).To(Succeed())
		})

		It("newLearningTrackingData returns error of creating scan results", func() {
			mockScanResultRep.EXPECT().CreateAll(gomock.Any()).Return(testErr).Times(1)
			Expect(manager.newLearningTrackingData(id, scanRes)).To(MatchError(testErr))
		})
	})

	Context("Test FinishLearning", func() {
		outRoom := &models.Room{ID: roomID}

		It("FinishLearning throws error starting with 'tracker' if tracker does not exist", func() {
			mockTrackerRep.EXPECT().GetByID(wrongID).Return(nil, recordNotFoundErr).Times(1)
			Expect(manager.FinishLearning(wrongID, roomID, []string{}).Error()).To(HavePrefix("tracker: "))
		})

		It("FinishLearning throws error if tracker has not status LearningFinished", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerIdle, nil).Times(1)
			Expect(manager.FinishLearning(id, roomID, []string{})).To(HaveOccurred())
		})

		It("FinishLearning throws error starting with 'room' if room does not exist", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearningFinished, nil).Times(1)
			mockRoomRep.EXPECT().GetByID(wrongRoomID).Return(nil, recordNotFoundErr).Times(1)
			Expect(manager.FinishLearning(id, wrongRoomID, []string{}).Error()).To(HavePrefix("room: "))
		})

		It("FinishLearning throws an error getting if scan results errors", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearningFinished, nil).Times(1)
			mockRoomRep.EXPECT().GetByID(roomID).Return(outRoom, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(id).Return(nil, recordNotFoundErr).Times(1)
			mockTemplateRep.EXPECT().GetStepsByRoomID(roomID).Return([]*models.Step{&models.Step{}}, nil).Times(1)
			Expect(manager.FinishLearning(id, roomID, []string{}).Error()).To(HavePrefix("scanResults: "))
		})

		It("FinishLearning sets room.IsLearned to true and tracker status to idle", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearningFinished, nil).Times(1)
			mockRoomRep.EXPECT().GetByID(roomID).Return(outRoom, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(id).Return([]*models.ScanResult{}, nil)
			mockRoomRep.EXPECT().Update(outRoom).Return(nil).Times(1)
			mockTrackerRep.EXPECT().SetStatusByID(id, models.TrackerStatusIdle).Return(nil).Times(1)
			mockTemplateRep.EXPECT().GetStepsByRoomID(roomID).Return([]*models.Step{&models.Step{}}, nil).Times(1)
			Expect(manager.FinishLearning(id, roomID, []string{})).To(Succeed())
		})
	})

	Context("Test GetLearningStatus", func() {
		It("GetLearningStatus throws error if tracker does not exist", func() {
			mockTrackerRep.EXPECT().GetByID(wrongID).Return(nil, recordNotFoundErr).Times(1)
			_, _, err := manager.GetLearningStatus(wrongID)
			Expect(err).To(MatchError(recordNotFoundErr))
		})

		It("GetLearningStatus throws error if tracker is not in learning or learningFinished", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerIdle, nil).Times(1)
			_, _, err := manager.GetLearningStatus(id)
			Expect(err).To(HaveOccurred())
		})

		It("GetLearningStatus return done is true if tracker status is learningFinished", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearningFinished, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(gomock.Any()).AnyTimes()
			done, _, _ := manager.GetLearningStatus(id)
			Expect(done).To(BeTrue())
		})

		It("GetLearningStatus return done is false if tracker status is learning", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearning, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(gomock.Any()).AnyTimes()
			done, _, _ := manager.GetLearningStatus(id)
			Expect(done).To(BeFalse())
		})

		It("GetLearningStatus returns correct ssid list", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearning, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(id).Return(scanResWithID, nil).Times(1)
			_, ssids, _ := manager.GetLearningStatus(id)
			Expect(ssids).To(ConsistOf(scanResSSIDs))
		})
	})
})
