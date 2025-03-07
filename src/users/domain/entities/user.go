package entities
import (
	"time"
)

type User struct {
    User_id        int  `json:"user_id"`
    Name      string    `json:"name" validate:"required,min=3,max=100"`
    Email     string    `json:"email" validate:"required,email"`
    Address   string    `json:"address" validate:"required,min=10,max=500"`
    CreatedAt time.Time `json:"created_at"`
}

func NewUser(user_id int, name string, email string, address string, created_at time.Time) *User {
	return &User{
		User_id:   user_id,
		Name:      name,
		Email:     email,
		Address:   address,
		CreatedAt: created_at,
	}
}

func (u *User) GetUserID() int {
	return u.User_id
}
func (u *User) GetName() string {
	return u.Name
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetAddress() string {
	return u.Address
}
func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}
