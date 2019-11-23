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
})
