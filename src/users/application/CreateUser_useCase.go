package application

import (
	"mi-tienda-online/src/users/domain"
	"mi-tienda-online/src/users/domain/entities"
)

type CreateUserUseCase struct {
	UserRepository domain.UserRepository
}

func NewCreateUserUseCase(UserRepository domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: UserRepository}
}

func (useCase *CreateUserUseCase) Execute(user *entities.User) error {
	err := useCase.UserRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
