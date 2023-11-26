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
//	@Param			RoomType	body		RoomTypeCreate	true	"RoomTypeCreate"
//	@Success		200			{object}	RoomTypeCreate
//	@Router			/room-types [post]
func (controller *RoomTypeController) CreateRoomType() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RoomTypeCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		if err := controller.biz.CreateRoomType(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(RoomTypeEntityName, err))
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
//	@Param			id			path		string			true	"RoomType ID"
//	@Param			RoomType	body		RoomTypeUpdate	true	"RoomTypeUpdate"
//	@Success		200			{object}	RoomTypeUpdate
//	@Router			/room-types/{id} [put]
func (controller *RoomTypeController) UpdateRoomType() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := bases.GetIdFromFakeId(c.Param("id"))

		var data RoomTypeUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := controller.biz.UpdateRoomType(c.Request.Context(), id, &data); err != nil {
			panic(common.ErrCannotUpdateEntity(RoomTypeEntityName, err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}

// ListRoomTypes godoc
//
//	@Summary		List room types
//	@Description	List room types
//	@Tags			room-types
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	RoomTypeResponse
//	@Router			/room-types [get]
func (controller *RoomTypeController) ListRoomTypes() gin.HandlerFunc {
	return func(c *gin.Context) {
		roomTypes, err := controller.biz.GetRoomTypes(c.Request.Context())

		if err != nil {
			panic(common.ErrCannotListEntity(RoomTypeEntityName, err))
		}

		var result = make([]RoomTypeResponse, len(roomTypes))
		for i, roomType := range roomTypes {
			result[i] = MapToRoomTypeResponse(&roomType)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

// GetRoomType godoc
//
//	@Summary		Get room type
//	@Description	Get room type
//	@Tags			room-types
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"RoomType ID"
//	@Success		200	{object}	RoomTypeResponse
//	@Router			/room-types/{id} [get]
func (controller *RoomTypeController) GetRoomType() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := bases.GetIdFromFakeId(c.Param("id"))

		roomType, err := controller.biz.GetRoomType(c.Request.Context(), id)

		if err != nil {
			panic(common.ErrCannotGetEntity(RoomTypeEntityName, err))
		}

		var result = MapToRoomTypeResponse(roomType)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
