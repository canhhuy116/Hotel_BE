package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/component"
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

		if err := controller.roomBiz.CreateRoom(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(RoomEntityName, err))
		}

		data.FakeId(common.DbRoom)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeID))
	}
}
