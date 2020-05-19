package managers

import (
	"errors"
	"paper-tracker/mock"
	"paper-tracker/models"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TrackerManager", func() {
	var (
		mockTrackerRep *mock.MockTrackerRepository
		mockCtrl       *gomock.Controller
		manager        *TrackerManagerImpl

		//trackerIdle             *models.Tracker
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
		manager = CreateTrackerManager(mockTrackerRep).(*TrackerManagerImpl)

		//trackerIdle = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusIdle}
		trackerLearningFinished = &models.Tracker{ID: id, Label: "New Tracker", Status: models.TrackerStatusLearningFinished}

		gormNotFound := func(err error) bool {
			return err == recordNotFoundErr
		}
		mockTrackerRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Test GetAllTrackers", func() {
		outTrackers := []*models.Tracker{{ID: 1, Label: "Tracker 1"}}

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

	Context("Test NewTrackingData", func() {
		It("NewTrackingData throws error for tracker with status LearningFinished", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerLearningFinished, nil).Times(1)
			Expect(manager.NewTrackingData(id, 1, 1, nil)).To(HaveOccurred())
		})

		/*It("NewTrackingData throws error for tracker with status Idle", func() {
			mockTrackerRep.EXPECT().GetByID(id).Return(trackerIdle, nil).Times(1)
			Expect(manager.NewTrackingData(id, nil)).To(HaveOccurred())
		})*/
	})

	Context("Test InWorkingHours", func() {
		mgr := &TrackerManagerImpl{}

		It("Returns that we are in working hours, if either the start or end time are < 0", func() {
			var inHours bool
			inHours, _ = mgr.inWorkingHours(time.Now(), -1, 0, false)
			Expect(inHours).To(Equal(true))
			inHours, _ = mgr.inWorkingHours(time.Now(), 0, -1, false)
			Expect(inHours).To(Equal(true))
			inHours, _ = mgr.inWorkingHours(time.Now(), -1, -1, false)
			Expect(inHours).To(Equal(true))
			inHours, _ = mgr.inWorkingHours(time.Now(), -1, -1, true)
			Expect(inHours).To(Equal(true))
		})

		It("Returns that we are in working hours, if we are", func() {
			var inHours bool
			inHours, _ = mgr.inWorkingHours(time.Date(2020, time.January, 1, 10, 0, 0, 0, time.Local), 9, 16, false)
			Expect(inHours).To(Equal(true))
			inHours, _ = mgr.inWorkingHours(time.Date(2020, time.January, 1, 15, 0, 0, 0, time.Local), 9, 16, false)
			Expect(inHours).To(Equal(true))
			inHours, _ = mgr.inWorkingHours(time.Date(2020, time.January, 1, 15, 59, 59, 0, time.Local), 9, 16, false)
			Expect(inHours).To(Equal(true))
		})

		It("Returns that we are not in working hours, if we aren't", func() {
			var inHours bool
			inHours, _ = mgr.inWorkingHours(time.Date(2020, time.January, 1, 8, 0, 0, 0, time.Local), 9, 16, false)
			Expect(inHours).To(Equal(false))
			inHours, _ = mgr.inWorkingHours(time.Date(2020, time.January, 1, 8, 59, 59, 0, time.Local), 9, 16, false)
			Expect(inHours).To(Equal(false))
			inHours, _ = mgr.inWorkingHours(time.Date(2020, time.January, 1, 16, 0, 0, 1, time.Local), 9, 16, false)
			Expect(inHours).To(Equal(false))
		})

	})

})
