package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jarqvi/gin-crud/models"
	"github.com/jarqvi/gin-crud/routes"
)

func main() {
	models.InitDB()
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running on port 8080.",
		})
	})

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
