package routes

import (
	"github.com/Harshraj13/go-jwt/controllers"
	"github.com/gin-gonic/gin"
)

// "github.com/Harshraj13/go-jwt/"

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("user/login", controllers.Login())
}
