package managers

import (
	"errors"
	"paper-tracker/mock"
	"paper-tracker/models"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RoomManager", func() {
	var (
		mockRoomRep     *mock.MockRoomRepository
		mockTemplateRep *mock.MockWorkflowTemplateRepository
		mockCtrl        *gomock.Controller
		manager         *RoomManagerImpl

		testErr = errors.New("error")
	)
	const (
		id models.RoomID = 1
	)

	BeforeEach(func() {
		roomManager = nil
		workflowTemplateManager = nil

		mockCtrl = gomock.NewController(GinkgoT())
		mockRoomRep = mock.NewMockRoomRepository(mockCtrl)
		mockTemplateRep = mock.NewMockWorkflowTemplateRepository(mockCtrl)
		manager = CreateRoomManager(mockRoomRep).(*RoomManagerImpl)
		CreateWorkflowTemplateManager(mockTemplateRep)

		gormNotFound := func(err error) bool {
			return gorm.IsRecordNotFoundError(err)
		}
		mockRoomRep.EXPECT().IsRecordNotFoundError(gomock.Any()).DoAndReturn(gormNotFound).AnyTimes()
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Test CreateRoom", func() {
		roomID0 := &models.Room{ID: 0, Label: "TestRoom"}
		roomID1 := &models.Room{ID: 1, Label: "TestRoom"}

		It("CreateRoom should call create in rep exactly once", func() {
			mockRoomRep.EXPECT().Create(roomID0).Return(nil).Times(1)
			Expect(manager.CreateRoom(roomID0)).To(Succeed())
		})

		It("CreateRoom should set given room id to 0", func() {
			mockRoomRep.EXPECT().Create(roomID0).Return(nil).Times(1)
			Expect(manager.CreateRoom(roomID1)).To(Succeed())
		})

		It("CreateRoom should return db error", func() {
			mockRoomRep.EXPECT().Create(roomID0).Return(testErr).Times(1)
			Expect(manager.CreateRoom(roomID0)).To(MatchError(testErr))
		})
	})

	Context("Test GetAllRooms", func() {
		outRooms := []*models.Room{&models.Room{ID: id, Label: "Test Room", DeleteLocked: true}}

		It("GetAllRooms should call get all in rep exactly once", func() {
			mockRoomRep.EXPECT().GetAll().Return(outRooms, nil).Times(1)
			mockTemplateRep.EXPECT().GetStepsByRoomID(id).Return([]*models.Step{&models.Step{}}, nil).Times(1)
			Expect(manager.GetAllRooms()).To(Equal(outRooms))
		})

		It("GetAllTrackers should return db error", func() {
			mockRoomRep.EXPECT().GetAll().Return(nil, testErr).Times(1)
			_, err := manager.GetAllRooms()
			Expect(err).To(MatchError(testErr))
		})
	})

	Context("Test GetRoomByID", func() {
		outRoom := &models.Room{ID: id, Label: "TestRoom"}

		It("GetRoomByID calls getByID in rep exactly once", func() {
			mockRoomRep.EXPECT().GetByID(id).Return(outRoom, nil).Times(1)
			mockTemplateRep.EXPECT().GetStepsByRoomID(id).Return([]*models.Step{&models.Step{}}, nil).Times(1)
			Expect(manager.GetRoomByID(id)).To(Equal(outRoom))
		})

		It("GetRoomByID should return db error", func() {
			mockRoomRep.EXPECT().GetByID(id).Return(nil, testErr).Times(1)
			_, err := manager.GetRoomByID(id)
			Expect(err).To(MatchError(testErr))
		})
	})

	Context("Test SetRoomLearned", func() {
		It("SetRoomLearned with learned true calls SetLearnedByID with learned true in rep exactly once", func() {
			mockRoomRep.EXPECT().SetLearnedByID(id, true).Return(nil).Times(1)
			Expect(manager.SetRoomLearned(id, true)).To(Succeed())
		})

		It("SetRoomLearned should return db error", func() {
			mockRoomRep.EXPECT().SetLearnedByID(id, true).Return(testErr).Times(1)
			Expect(manager.SetRoomLearned(id, true)).To(MatchError(testErr))
		})

		It("SetRoomLearned also works for learned false", func() {
			mockRoomRep.EXPECT().SetLearnedByID(id, false).Return(nil).Times(1)
			Expect(manager.SetRoomLearned(id, false)).To(Succeed())
		})
	})

	Context("Test UpdateRoom", func() {
		room := &models.Room{ID: id, Label: "TestRoom"}

		It("UpdateRoom calls rep update exactly once", func() {
			mockRoomRep.EXPECT().Update(room).Return(nil).Times(1)
			Expect(manager.UpdateRoom(room)).To(Succeed())
		})

		It("UpdateRoom should return db error", func() {
			mockRoomRep.EXPECT().Update(room).Return(testErr).Times(1)
			Expect(manager.UpdateRoom(room)).To(MatchError(testErr))
		})
	})

	Context("Test DeleteRoom", func() {
		It("DeleteRoom calls rep delete exactly once", func() {
			mockRoomRep.EXPECT().Delete(id).Return(nil).Times(1)
			Expect(manager.DeleteRoom(id)).To(Succeed())
		})

		It("DeleteRoom should return db error", func() {
			mockRoomRep.EXPECT().Delete(id).Return(testErr).Times(1)
			Expect(manager.DeleteRoom(id)).To(MatchError(testErr))
		})
	})
})
