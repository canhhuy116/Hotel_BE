package users

import (
	"Hotel_BE/common"
	"Hotel_BE/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// CreateUser godoc
//
//	@Summary		Create user
//	@Description	create user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		UserCreate	true	"UserCreate"
//	@Success		200		{object}	UserCreate
//	@Router			/users [post]
func (c *UserController) CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		store := NewUserStore(appCtx.GetMainDbConnection())
		biz := NewUserBiz(store)

		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(EntityName, err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

// ListUser godoc
//
//	@Summary		List users
//	@Description	Get a list of users with pagination
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			page	query	int		false	"Page number (default is 1)"
//	@Param			limit	query	int		false	"Number of items per page (default is 10)"
//	@Param			cursor	query	string	false	"Cursor for pagination"
//	@Success		200		{array}	User
//	@Router			/users [get]
func (c *UserController) ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()

		store := NewUserStore(appCtx.GetMainDbConnection())
		biz := NewUserBiz(store)

		users, err := biz.ListUser(c.Request.Context(), &paging)

		if err != nil {
			panic(common.ErrCannotListEntity(EntityName, err))
		}

		for i := range users {
			users[i].Mask()

			if i == len(users)-1 {
				paging.NextCursor = users[i].FakeID
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, nil))
	}
}
