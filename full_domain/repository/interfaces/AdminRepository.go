package interfaces

import "full_domain/entity"

type AdminRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	FindUserById(id int) (*entity.User, error)
	DeleteUser(id int) (*entity.User, error)
	UpdateUser(id int, user *entity.User) (*entity.User, error)
	FindUserContaining(str string) ([]*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	FindAllusers() ([]*entity.User, error)
	BlockUser(id int) (*entity.User, error)
	UnBlockUser(id int) (*entity.User, error)
}
