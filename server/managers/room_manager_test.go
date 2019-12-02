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
		mockRoomRep *mock.MockRoomRepository
		mockCtrl    *gomock.Controller
		manager     *RoomManager

		testErr = errors.New("error")
	)

	BeforeEach(func() {
		roomManager = nil

		mockCtrl = gomock.NewController(GinkgoT())
		mockRoomRep = mock.NewMockRoomRepository(mockCtrl)
		manager = CreateRoomManager(mockRoomRep)

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
		outRooms := []*models.Room{&models.Room{Label: "Test Room"}}

		It("GetAllRooms should call get all in rep exactly once", func() {
			mockRoomRep.EXPECT().GetAll().Return(outRooms, nil).Times(1)
			Expect(manager.GetAllRooms()).To(Equal(outRooms))
		})

		It("GetAllTrackers should return db error", func() {
			mockRoomRep.EXPECT().GetAll().Return(nil, testErr).Times(1)
			_, err := manager.GetAllRooms()
			Expect(err).To(MatchError(testErr))
		})
	})

	Context("Test GetRoomByID", func() {
		id := 1
		outRoom := &models.Room{ID: id, Label: "TestRoom"}

		It("GetRoomByID calls getByID in rep exactly once", func() {
			mockRoomRep.EXPECT().GetByID(id).Return(outRoom, nil).Times(1)
			Expect(manager.GetRoomByID(id)).To(Equal(outRoom))
		})

		It("GetRoomByID should return db error", func() {
			mockRoomRep.EXPECT().GetByID(id).Return(nil, testErr).Times(1)
			_, err := manager.GetRoomByID(id)
			Expect(err).To(MatchError(testErr))
		})
	})
})
