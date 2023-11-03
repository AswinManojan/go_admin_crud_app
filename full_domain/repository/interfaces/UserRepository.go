package interfaces

import "full_domain/entity"

type UserRepository interface {
	RegisterUser(user *entity.User) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
}
