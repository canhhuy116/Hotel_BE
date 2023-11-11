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
// @Summary      Create user
// @Description  create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body UserCreate true "UserCreate"
// @Success      200  {object} UserCreate
// @Router       /users [post]
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

		c.JSON(http.StatusOK, gin.H{"user": common.SimpleSuccessResponse(data)})
	}
}

// ListUser  godoc
// @Summary      List users
// @Description  get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   User
// @Router       /users [get]
func (c *UserController) ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := NewUserStore(appCtx.GetMainDbConnection())
		biz := NewUserBiz(store)

		users, err := biz.ListUser(c.Request.Context())

		if err != nil {
			panic(common.ErrCannotListEntity(EntityName, err))
		}

		for i := range users {
			users[i].Mask()
		}

		c.JSON(http.StatusOK, gin.H{"users": common.SimpleSuccessResponse(users)})
	}
}
