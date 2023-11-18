package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoomTypeController struct{}

func NewRoomTypeController() *RoomTypeController {
	return &RoomTypeController{}
}

// CreateRoomType godoc
//
//	@Summary		Create room type
//	@Description	Create room type
//	@Tags			room-types
//	@Accept			json
//	@Produce		json
//	@Param			RoomType	body		RoomTypeCreate	true "RoomTypeCreate"
//	@Success		200			{object}	RoomTypeCreate
//	@Router			/room-types [post]
func (c *RoomTypeController) CreateRoomType(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RoomTypeCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		store := NewRoomTypeStore(appCtx.GetMainDbConnection())
		biz := NewRoomTypeBiz(store)

		if err := biz.CreateRoomType(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(EntityName, err))
		}

		data.FakeId(common.DbRoomType)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeID))
	}
}
