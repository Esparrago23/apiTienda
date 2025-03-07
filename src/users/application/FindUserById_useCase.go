package application

import (
	"mi-tienda-online/src/users/domain"
	"mi-tienda-online/src/users/domain/entities"
)

type FindUserByIdUseCase struct {
	UserRepository domain.UserRepository
}

func NewFindUserByIdUseCase(UserRepository domain.UserRepository) *FindUserByIdUseCase {
	return &FindUserByIdUseCase{UserRepository: UserRepository}
}

func (useCase *FindUserByIdUseCase) Execute(user_id int) (*entities.User, error) {
	user, err := useCase.UserRepository.FindById(user_id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
