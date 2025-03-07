package application

import (
	"mi-tienda-online/src/users/domain"
	"mi-tienda-online/src/users/domain/entities"
)

type UpdateUserUseCase struct {
	UserRepository domain.UserRepository
}

func NewUpdateUserUseCase(UserRepository domain.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserRepository: UserRepository}
}

func (useCase *UpdateUserUseCase) Execute(user *entities.User) error {
	err := useCase.UserRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}
