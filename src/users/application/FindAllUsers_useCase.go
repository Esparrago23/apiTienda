package application

import (
	"mi-tienda-online/src/users/domain"
	"mi-tienda-online/src/users/domain/entities"
)

type FindAllUsersUseCase struct {
	UserRepository domain.UserRepository
}

func NewFindAllUsersUseCase(UserRepository domain.UserRepository) *FindAllUsersUseCase {
	return &FindAllUsersUseCase{UserRepository: UserRepository}
}

func (useCase *FindAllUsersUseCase) Execute() ([]entities.User, error) {
	users, err := useCase.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
