package interfaces

import (
	"full_domain/dto"
	"full_domain/entity"
)

type AdminService interface {
	Login(loginRequest *dto.LoginRequest) (string, error)
	DeleteUser(id int) (*entity.User, error)
	UpdateUser(id int, user entity.User) (*entity.User, error)
	BlockUser(id int) (*entity.User, error)
	UnBlockUser(id int) (*entity.User, error)
	AddUser(user *entity.User) (*entity.User, error)
	FindUser(id int) (*entity.User, error)
	SearchUser(str string) ([]*entity.User, error)
	FindAllUsers() ([]*entity.User, error)
}
