package main

import (
	"boards-api/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api")

	createUserRoutes(api)

	err := router.Run()
	if err != nil {
		panic(err)

	}
}

func createUserRoutes(group *gin.RouterGroup) {
	user := group.Group("/user")

	userv1 := user.Group("/v1")
	userv1.POST("login", http.CreateUserHandler)
}
