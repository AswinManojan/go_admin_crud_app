package service

import (
	"errors"
	"fmt"
	"full_domain/dto"
	"full_domain/entity"
	"full_domain/middleware"
	repository "full_domain/repository/interfaces"
	service "full_domain/service/interfaces"
	"log"
)

type AdminServiceImpl struct {
	repo repository.AdminRepository
	jwt  *middleware.JwtUtil
}

// AddUser implements interfaces.AdminService.
func (asi *AdminServiceImpl) AddUser(user *entity.User) (*entity.User, error) {
	user, err := asi.repo.CreateUser(user)
	if err != nil {
		log.Println("User not added, adminService file")
		return user, err
	}
	return user, nil
}

// BlockUser implements interfaces.AdminService.
func (asi *AdminServiceImpl) BlockUser(id int) (*entity.User, error) {
	user, err := asi.repo.BlockUser(id)
	if err != nil {
		log.Println("Error Blocking user, in adminServiceImpl file")
		return user, err
	}
	return user, err
}

// DeleteUser implements interfaces.AdminService.
func (asi *AdminServiceImpl) DeleteUser(id int) (*entity.User, error) {
	fmt.Print(id)
	user, err := asi.repo.DeleteUser(id)
	if err != nil {
		log.Println("Error Deleting user, in adminServiceImpl file")
		return user, err
	}
	return user, err
}

// FindAllUsers implements interfaces.AdminService.
func (asi *AdminServiceImpl) FindAllUsers() ([]*entity.User, error) {
	userList, err := asi.repo.FindAllusers()
	if err != nil {
		log.Println("Error Finding All users, in adminServiceImpl file")
		return userList, err
	}
	return userList, err
}

// FindUser implements interfaces.AdminService.
func (asi *AdminServiceImpl) FindUser(id int) (*entity.User, error) {
	user, err := asi.repo.FindUserById(id)
	if err != nil {
		log.Println("Error Finding user, in adminServiceImpl file")
		return user, err
	}

	return user, err
}

// Login implements interfaces.AdminService.
func (asi *AdminServiceImpl) Login(loginRequest *dto.LoginRequest) (string, error) {
	user, err := asi.repo.FindUserByEmail(loginRequest.Email)
	if err != nil {
		log.Println("No USER EXISTS, in adminService file")
		return "", errors.New("no User exists")
	}

	if user.Password != loginRequest.Password {
		log.Println("Password Mismatch, in adminService file")
		return "", errors.New("password Mismatch")
	}

	if user.Role != "admin" {
		log.Println("Unauthorized, in adminService file")
		return "", errors.New("unauthorized access")
	}

	token, err := asi.jwt.CreateToken(loginRequest.Email, "admin")
	if err != nil {
		return "", errors.New("token NOT generated")
	}

	return token, nil
}

// SearchUser implements interfaces.AdminService.
func (asi *AdminServiceImpl) SearchUser(str string) ([]*entity.User, error) {
	userList, err := asi.repo.FindUserContaining(str)
	if err != nil {
		log.Println("Error Searching users, in adminServiceImpl file")
		return userList, err
	}

	return userList, err
}

// UnBlockUser implements interfaces.AdminService.
func (asi *AdminServiceImpl) UnBlockUser(id int) (*entity.User, error) {
	user, err := asi.repo.UnBlockUser(id)
	if err != nil {
		log.Println("Error Blocking user, in adminServiceImpl file")
		return user, err
	}
	return user, err
}

// UpdateUser implements interfaces.AdminService.
func (asi *AdminServiceImpl) UpdateUser(id int, user entity.User) (*entity.User, error) {
	updatedUser, err := asi.repo.UpdateUser(id, &user)

	if err != nil {
		log.Println("Error Updating user, in adminServiceImpl file")
		return updatedUser, err
	}

	return updatedUser, nil
}

func NewAdminService(repository repository.AdminRepository, jwt *middleware.JwtUtil) service.AdminService {
	return &AdminServiceImpl{
		repo: repository,
		jwt:  jwt,
	}
}
