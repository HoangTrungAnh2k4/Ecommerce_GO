package controller

import (
	"github.com/HoangTrungAnh2k4/Ecommerce_GO/internal/service"
	"github.com/HoangTrungAnh2k4/Ecommerce_GO/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc: user controller
func (uc *UserController) RegisterUser(c *gin.Context) {
	response.SuccessResponse(c, 2001, []string{"ok"})
}
