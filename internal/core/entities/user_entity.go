package entities

import (
	"boards-api/internal/core/entities/enums"

	"github.com/samborkent/uuidv7"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	ID       uuidv7.UUID    `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Role     enums.RoleEnum `json:"role"`
}

func (ue *UserEntity) ComparePassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ue.Password), []byte(plain))
	return err == nil
}
