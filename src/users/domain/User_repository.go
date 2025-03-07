package domain

import (
	"mi-tienda-online/src/users/domain/entities"
)

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindById(user_id int) (entities.User, error)
	Save(user *entities.User) error
	Update(user *entities.User) error
	Delete(user_id int) error
}