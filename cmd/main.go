package main

import (
	_ "boards-api/docs"
	"boards-api/internal/adapters/http"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title						TODO APIs
// @version					1.0
// @description				Testing Swagger APIs.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						token
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8081
// @BasePath					/api/v1
// @schemes					http
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api")

	createUserRoutes(api)

	err = router.Run()
	if err != nil {
		panic(err)

	}
}

func createUserRoutes(group *gin.RouterGroup) {
	user := group.Group("user")

	userv1 := user.Group("v1")

	userv1.POST("register", http.CreateUserHandler)
}
