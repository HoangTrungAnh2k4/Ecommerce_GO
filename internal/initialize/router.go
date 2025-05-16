package initialize

import (
	"github.com/HoangTrungAnh2k4/Ecommerce_GO/internal/controller"
	"github.com/HoangTrungAnh2k4/Ecommerce_GO/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewUserController().RegisterUser)
	}
	return r
}
