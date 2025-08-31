package ports

import (
	"boards-api/internal/adapters/repositories"
	"boards-api/internal/core/entities"
)

type SaveUserInput struct {
	User entities.UserEntity
}

type SaveUserPort interface {
	SaveUser(input SaveUserInput) error
}

type saveUser struct {
	repo repositories.UserRepository
}

func NewSaveUserPort(repo repositories.UserRepository) SaveUserPort {
	return &saveUser{
		repo: repo,
	}
}

func (s saveUser) SaveUser(input SaveUserInput) error {
	err := s.repo.SaveUser(input.User)
	if err != nil {
		return err
	}
	return nil
}
