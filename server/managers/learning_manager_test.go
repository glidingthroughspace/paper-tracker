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
		mockScanResultRep *mock.MockScanResultRepository
		mockCtrl          *gomock.Controller
		manager           *LearningManagerImpl

		mockTrackerManager          *mock.MockTrackerManager
		mockRoomManager             *mock.MockRoomManager
		mockWorkflowTemplateManager *mock.MockWorkflowTemplateManager
		mockTrackingManager         *mock.MockTrackingManager

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

		mockCtrl = gomock.NewController(GinkgoT())
		mockScanResultRep = mock.NewMockScanResultRepository(mockCtrl)
		manager = CreateLearningManager(mockScanResultRep).(*LearningManagerImpl)

		mockTrackerManager = mock.NewMockTrackerManager(mockCtrl)
		trackerManager = mockTrackerManager
		mockRoomManager = mock.NewMockRoomManager(mockCtrl)
		roomManager = mockRoomManager
		mockWorkflowTemplateManager = mock.NewMockWorkflowTemplateManager(mockCtrl)
		workflowTemplateManager = mockWorkflowTemplateManager
		mockTrackingManager = mock.NewMockTrackingManager(mockCtrl)
		trackingManager = mockTrackingManager

		gormNotFound := func(err error) bool {
			return gorm.IsRecordNotFoundError(err)
		}
		mockScanResultRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()

		trackerIdle = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusIdle}
		trackerLearning = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusLearning}
		trackerLearningFinished = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusLearningFinished}

		scanRes = []*models.ScanResult{
			{SSID: "Test0", BSSID: "aa:bb:cc:dd:ee:ff", RSSI: -50},
			{SSID: "Test1", BSSID: "aa:bb:cc:dd:ee:ff", RSSI: -40},
		}
		scanResWithID = []*models.ScanResult{
			{TrackerID: id, SSID: "Test0", BSSID: "aa:bb:cc:dd:ee:ff", RSSI: -50},
			{TrackerID: id, SSID: "Test1", BSSID: "aa:bb:cc:dd:ee:ff", RSSI: -40},
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
			trackerSetStatusCall = mockTrackerManager.EXPECT().SetTrackerStatus(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
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
				mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerLearning, nil).AnyTimes()
				manager.learningRoutine(id, testLogger)
			})
		})
	})

	Context("Test NewLearningTrackingData", func() {
		It("NewLearningTrackingData creates all scan results in db", func() {
			mockScanResultRep.EXPECT().CreateAll(gomock.Any()).Return(nil).Times(1)
			Expect(manager.NewLearningTrackingData(id, scanRes)).To(Succeed())
		})

		It("NewLearningTrackingData add proper trackerID to scan results", func() {
			mockScanResultRep.EXPECT().CreateAll(scanResWithID).Return(nil).Times(1)
			Expect(manager.NewLearningTrackingData(id, scanRes)).To(Succeed())
		})

		It("NewLearningTrackingData returns error of creating scan results", func() {
			mockScanResultRep.EXPECT().CreateAll(gomock.Any()).Return(testErr).Times(1)
			Expect(manager.NewLearningTrackingData(id, scanRes)).To(MatchError(testErr))
		})
	})

	Context("Test FinishLearning", func() {
		outRoom := &models.Room{ID: roomID}

		It("FinishLearning throws error starting with 'tracker' if tracker does not exist", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(wrongID).Return(nil, recordNotFoundErr).Times(1)
			Expect(manager.FinishLearning(wrongID, roomID, []string{}).Error()).To(HavePrefix("tracker: "))
		})

		It("FinishLearning throws error if tracker has not status LearningFinished", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerIdle, nil).Times(1)
			Expect(manager.FinishLearning(id, roomID, []string{})).To(HaveOccurred())
		})

		It("FinishLearning throws error starting with 'room' if room does not exist", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerLearningFinished, nil).Times(1)
			mockRoomManager.EXPECT().GetRoomByID(wrongRoomID).Return(nil, recordNotFoundErr).Times(1)
			Expect(manager.FinishLearning(id, wrongRoomID, []string{}).Error()).To(HavePrefix("room: "))
		})

		It("FinishLearning throws an error getting if scan results errors", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerLearningFinished, nil).Times(1)
			mockRoomManager.EXPECT().GetRoomByID(roomID).Return(outRoom, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(id).Return(nil, recordNotFoundErr).Times(1)
			Expect(manager.FinishLearning(id, roomID, []string{}).Error()).To(HavePrefix("scanResults: "))
		})

		It("FinishLearning sets room.IsLearned to true and tracker status to idle", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerLearningFinished, nil).Times(1)
			mockRoomManager.EXPECT().GetRoomByID(roomID).Return(outRoom, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(id).Return([]*models.ScanResult{}, nil)
			mockRoomManager.EXPECT().UpdateRoom(outRoom).Return(nil).Times(1)
			mockTrackerManager.EXPECT().SetTrackerStatus(id, models.TrackerStatusIdle).Return(nil).Times(1)
			mockTrackingManager.EXPECT().ConsolidateScanResults([]*models.ScanResult{}).Return([]models.BSSIDTrackingData{}).Times(1)
			Expect(manager.FinishLearning(id, roomID, []string{})).To(Succeed())
		})
	})

	Context("Test GetLearningStatus", func() {
		It("GetLearningStatus throws error if tracker does not exist", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(wrongID).Return(nil, recordNotFoundErr).Times(1)
			_, _, err := manager.GetLearningStatus(wrongID)
			Expect(err).To(MatchError(recordNotFoundErr))
		})

		It("GetLearningStatus throws error if tracker is not in learning or learningFinished", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerIdle, nil).Times(1)
			_, _, err := manager.GetLearningStatus(id)
			Expect(err).To(HaveOccurred())
		})

		It("GetLearningStatus return done is true if tracker status is learningFinished", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerLearningFinished, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(gomock.Any()).AnyTimes()
			done, _, _ := manager.GetLearningStatus(id)
			Expect(done).To(BeTrue())
		})

		It("GetLearningStatus return done is false if tracker status is learning", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerLearning, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(gomock.Any()).AnyTimes()
			done, _, _ := manager.GetLearningStatus(id)
			Expect(done).To(BeFalse())
		})

		It("GetLearningStatus returns correct ssid list", func() {
			mockTrackerManager.EXPECT().GetTrackerByID(id).Return(trackerLearning, nil).Times(1)
			mockScanResultRep.EXPECT().GetAllForTracker(id).Return(scanResWithID, nil).Times(1)
			_, ssids, _ := manager.GetLearningStatus(id)
			Expect(ssids).To(ConsistOf(scanResSSIDs))
		})
	})
})
