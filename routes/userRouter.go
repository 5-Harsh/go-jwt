package routes

import (
	controller "github.com/Harshraj13/go-jwt/controllers"
	"github.com/Harshraj13/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
