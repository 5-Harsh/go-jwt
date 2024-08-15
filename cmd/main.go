package main

import (

	// routes "github.com/Harshraj13/go-jwt/routes"

	"github.com/Harshraj13/go-jwt/database"
	"github.com/Harshraj13/go-jwt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := database.LoadEnvVariable("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// router.GET("/api-1", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"success": "Acess granted for api-1"})
	// })

	// router.GET("/api-2", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"success": "Access granted for api-2"})
	// })

	router.Run(":" + port)

}
