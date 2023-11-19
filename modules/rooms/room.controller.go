package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/component"
	"Hotel_BE/modules/bases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoomController struct {
	roomBiz *RoomBiz
}

func NewRoomController(appCtx component.AppContext) *RoomController {
	store := NewRoomStore(appCtx.GetMainDbConnection())
	biz := NewRoomBiz(store)
	return &RoomController{roomBiz: biz}
}

// CreateRoom godoc
//
//	@Summary		Create room
//	@Description	Create room
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param			Room	body		RoomCreate	true	"RoomCreate"
//	@Success		200		{object}	RoomCreate
//	@Router			/rooms [post]
func (controller *RoomController) CreateRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RoomCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		data.RoomTypeId = bases.GetIdFromFakeId(data.RoomTypeFakeId)

		if err := controller.roomBiz.CreateRoom(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(RoomEntityName, err))
		}

		data.FakeId(common.DbRoom)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeID))
	}
}

// UpdateRoom godoc
//
//	@Summary		Update room
//	@Description	Update room
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string		true	"Room ID"
//	@Param			Room	body	RoomUpdate	true	"RoomUpdate"
//	@Success		200
//	@Router			/rooms/{id} [put]
func (controller *RoomController) UpdateRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := bases.GetIdFromFakeId(c.Param("id"))

		var data RoomUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		if err := controller.roomBiz.UpdateRoom(c.Request.Context(), id, &data); err != nil {
			panic(common.ErrCannotUpdateEntity(RoomEntityName, err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
