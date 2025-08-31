package http

import (
	"boards-api/internal/adapters/repositories"
	"boards-api/internal/core/usecases/createuser"
	"boards-api/internal/core/usecases/createuser/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var dto createuser.CreateUserInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     err.Error(),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	repo := repositories.NewUserRepository()
	port := ports.NewSaveUserPort(repo)
	usecase := createuser.NewCreateUserUseCase(port)

	user, err := usecase.CreateUser(&dto)
	if err != nil {
		return
	}
	c.JSON(user.StatusCode, user)
}
