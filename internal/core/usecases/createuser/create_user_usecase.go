package createuser

import (
	"boards-api/internal/core/entities"
	"boards-api/internal/core/entities/enums"
	"boards-api/internal/core/usecases/createuser/ports"

	"github.com/samborkent/uuidv7"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name     string         `json:"name" binding:"required"`
	Email    string         `json:"email" binding:"required,email"`
	Password string         `json:"password" binding:"required,min=12"`
	Role     enums.RoleEnum `json:"role" binding:"required"`
}

type CreateUserOutput struct {
	Message    string `json:"message"`
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
}

type CreateUserUseCase interface {
	CreateUser(input *CreateUserInput) (*CreateUserOutput, error)
}

type createUser struct {
	savePort ports.SaveUserPort
}

func NewCreateUserUseCase(savePort ports.SaveUserPort) CreateUserUseCase {
	return &createUser{
		savePort: savePort,
	}
}

func (c createUser) CreateUser(input *CreateUserInput) (*CreateUserOutput, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	user := entities.UserEntity{
		ID:       uuidv7.New(),
		Name:     input.Name,
		Email:    input.Email,
		Role:     input.Role,
		Password: string(hash),
	}

	err = c.savePort.SaveUser(ports.SaveUserInput{
		User: user,
	})

	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		Message:    "Created.",
		StatusCode: 201,
		Path:       "/api/user/v1/login",
	}, nil
}
