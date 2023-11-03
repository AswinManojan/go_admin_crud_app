package interfaces

import (
	"full_domain/dto"
	"full_domain/entity"
)

type UserService interface {
	Login(loginRequest *dto.LoginRequest) (string, error)
	RegisterUser(user *entity.User) (*entity.User, error)
}
