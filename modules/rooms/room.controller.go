package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/component"
	"Hotel_BE/modules/bases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomBiz *RoomBiz
}

func NewRoomController(appCtx component.AppContext) *RoomController {
	store := NewRoomStore(appCtx.GetMainDbConnection())
	biz := NewRoomBiz(store)
	return &RoomController{roomBiz: biz}
}

// GetRooms godoc
//
//	@Summary		Get rooms
//	@Description	Get rooms
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int		false	"Page"
//	@Param			limit	query		int		false	"Limit"
//	@Param			cursor	query		string	false	"Cursor"
//	@Success		200		{object}	Room
//	@Router			/rooms [get]
func (controller *RoomController) GetRooms() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBindQuery(&paging); err != nil {
			panic(err)
		}

		paging.FullFill()

		rooms, err := controller.roomBiz.GetRooms(c.Request.Context(), &paging)

		if err != nil {
			panic(common.ErrCannotGetEntity(RoomEntityName, err))
		}

		for i := range rooms {
			rooms[i].Mask()

			if i == len(rooms)-1 {
				paging.NextCursor = rooms[i].FakeID
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(rooms, paging, nil))
	}
}

// CreateRoom godoc
//
//	@Summary		Create room
//	@Description	Create room
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param			Room	body	RoomCreate	true	"RoomCreate"
//	@Success		200		{array}	RoomCreate
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

// GetRoom godoc
//
//	@Summary		Get room
//	@Description	Get room
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Room ID"
//	@Success		200	{object}	Room
//	@Router			/rooms/{id} [get]
func (controller *RoomController) GetRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := bases.GetIdFromFakeId(c.Param("id"))

		result, err := controller.roomBiz.GetRoom(c.Request.Context(), id)

		if err != nil {
			panic(common.ErrCannotGetEntity(RoomEntityName, err))
		}

		result.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
