package repository

import (
	"errors"
	"full_domain/entity"
	"full_domain/repository/interfaces"
	"log"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// FindUserByEmail implements interfaces.UserRepository.
func (uri *UserRepositoryImpl) FindUserByEmail(email string) (*entity.User, error) {
	if uri.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}

	user := &entity.User{}
	result := uri.DB.Where("email = ?", email).First(user)

	if result.Error != nil {
		log.Println("User not found in DB")
		return nil, errors.New("user not Found in DB")
	}

	return user, nil
}

// RegisterUser implements interfaces.UserRepository.
func (uri *UserRepositoryImpl) RegisterUser(user *entity.User) (*entity.User, error) {
	if uri.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return nil, errors.New("error Coonecting Database")
	}

	result := uri.DB.Create(&user)
	if result.Error != nil {
		log.Println("Unable to add user, AdminRepositoryImpl package")
		return user, errors.New("user already exists")
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
