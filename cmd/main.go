package main

import (
	"boards-api/internal/adapters/http"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
)

func main() {
	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	awsRegion := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithClientLogMode(aws.LogRequest|aws.LogRetries))

	if err != nil {
		panic(err)
	}

	cfg.BaseEndpoint = &awsEndpoint
	cfg.Region = awsRegion

	router := gin.Default()

	api := router.Group("/api")

	createUserRoutes(api)

	err = router.Run()
	if err != nil {
		panic(err)

	}
}

func createUserRoutes(group *gin.RouterGroup) {
	user := group.Group("/user")

	userv1 := user.Group("/v1")
	userv1.POST("login", http.CreateUserHandler)
}
