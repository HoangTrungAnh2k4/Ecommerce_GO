package middlewares

import (
	"github.com/HoangTrungAnh2k4/Ecommerce_GO/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrCodeInvalidToken, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
