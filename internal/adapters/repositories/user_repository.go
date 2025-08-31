package repositories

import "boards-api/internal/core/entities"

type UserRepository interface {
	SaveUser(entity entities.UserEntity) error
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (s *userRepository) SaveUser(entity entities.UserEntity) error {
	panic("implement me")
}
