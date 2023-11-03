package service

import (
	"errors"
	"full_domain/dto"
	"full_domain/entity"
	repository "full_domain/repository/interfaces"
	service "full_domain/service/interfaces"
	"log"

	"full_domain/middleware"
)

type UserServiceImpl struct {
	repo repository.UserRepository
	jwt  *middleware.JwtUtil
}

// Login implements interfaces.UserService.
func (usi *UserServiceImpl) Login(loginRequest *dto.LoginRequest) (string, error) {
	user, err := usi.repo.FindUserByEmail(loginRequest.Email)
	if err != nil {
		log.Println("No USER EXISTS, in adminService file")
		return "", errors.New("no User exists")
	}

	if user.Password != loginRequest.Password {
		log.Println("Password Mismatch, in adminService file")
		return "", errors.New("password Mismatch")
	}

	if user.Role != "user" {
		log.Println("Unauthorized, in adminService file")
		return "", errors.New("unauthorized access")
	}

	token, err := usi.jwt.CreateToken(loginRequest.Email, "user")
	if err != nil {
		return "", errors.New("token NOT generated")
	}
	return token, nil
}

// RegisterUser implements interfaces.UserService.
func (usi *UserServiceImpl) RegisterUser(user *entity.User) (*entity.User, error) {
	user, err := usi.repo.RegisterUser(user)
	if err != nil {
		log.Println("User not added, adminService file")
		return user, err
	}
	return user, err
}

func NewUserService(repository repository.UserRepository, jwt *middleware.JwtUtil) service.UserService {
	return &UserServiceImpl{
		repo: repository,
		jwt:  jwt,
	}
}
