package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/component"
	"Hotel_BE/modules/bases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoomTypeController struct {
	biz *RoomTypeBiz
}

func NewRoomTypeController(appCtx component.AppContext) *RoomTypeController {
	store := NewRoomTypeStore(appCtx.GetMainDbConnection())
	biz := NewRoomTypeBiz(store)
	return &RoomTypeController{biz: biz}
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
func (controller *RoomTypeController) CreateRoomType() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RoomTypeCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		if err := controller.biz.CreateRoomType(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(EntityName, err))
		}

		data.FakeId(common.DbRoomType)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeID))
	}
}

// UpdateRoomType godoc
//
//	@Summary		Update room type
//	@Description	Update room type
//	@Tags			room-types
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string		true	"RoomType ID"
//	@Param			RoomType	body	RoomTypeUpdate	true	"RoomTypeUpdate"
//	@Success		200		{object}	RoomTypeUpdate
//	@Router			/room-types/{id} [put]
func (controller *RoomTypeController) UpdateRoomType() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := bases.GetIdFromFakeId(c.Param("id"))

		var data RoomTypeUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := controller.biz.UpdateRoomType(c.Request.Context(), id, &data); err != nil {
			panic(common.ErrCannotUpdateEntity(EntityName, err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
