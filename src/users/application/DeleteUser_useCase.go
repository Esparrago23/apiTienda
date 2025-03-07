package application

import (
	"mi-tienda-online/src/users/domain"
)

type DeleteUserUseCase struct {
	UserRepository domain.UserRepository
}

func NewDeleteUserUseCase(UserRepository domain.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{UserRepository: UserRepository}
}

func (useCase *DeleteUserUseCase) Execute(user_id int) error {
	err := useCase.UserRepository.Delete(user_id)
	if err != nil {
		return err
	}
	return nil
}
