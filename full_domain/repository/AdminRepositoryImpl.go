package repository

import (
	"errors"
	"full_domain/entity"
	"full_domain/repository/interfaces"
	"log"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

// BlockUser implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) BlockUser(id int) (*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user, err := ari.FindUserById(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	user.IsLocked = true
	result := ari.DB.Save(user)
	if result.Error != nil {
		log.Println("Unable to block user")
		return nil, errors.New("unable to block user")
	}
	return user, nil
}

// CreateUser implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) CreateUser(user *entity.User) (*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}

	_, err := ari.FindUserByEmail(user.Email)
	if err == nil {
		log.Println("USER ALREADY EXISTS")
		return nil, errors.New("user exists in db")
	}

	result := ari.DB.Create(&user)
	if result.Error != nil {
		log.Println("Unable to add user, AdminRepositoryImpl package")
		return nil, errors.New("user not added to db")
	}
	return user, nil
}

// DeleteUser implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) DeleteUser(id int) (*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user, err := ari.FindUserById(id)
	if err != nil {
		log.Println("User Not found")
		return nil, errors.New("error deleting the user")
	}
	ari.DB.Delete(user).Where("id=?", id)
	// ari.DB.Raw("delete from users where id=1")
	return user, nil
}

// FindAllusers implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) FindAllusers() ([]*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	users := []*entity.User{}
	ari.DB.Find(&users)
	return users, nil
}

// FindUserByEmail implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) FindUserByEmail(email string) (*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user := &entity.User{}
	result := ari.DB.Where("email=?", email).First(user)
	if result.Error != nil {
		log.Println("User doesn't exist")
		return nil, errors.New("no user found for this email")
	}
	return user, nil
}

// FindUserById implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) FindUserById(id int) (*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user := &entity.User{}
	result := ari.DB.First(user,id)
	if result.Error != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	return user, nil
}

// FindUserContaining implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) FindUserContaining(str string) ([]*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}

	users := []*entity.User{}
	result := ari.DB.Where("email LIKE ?", "%"+str+"%").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// UnBlockUser implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) UnBlockUser(id int) (*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user, err := ari.FindUserById(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	user.IsLocked = false
	result := ari.DB.Save(user)
	if result.Error != nil {
		log.Println("Unable to block user")
		return nil, errors.New("unable to block user")
	}
	return user, nil
}

// UpdateUser implements interfaces.AdminRepository.
func (ari *AdminRepositoryImpl) UpdateUser(id int, user *entity.User) (*entity.User, error) {
	if ari.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	founduser, err := ari.FindUserById(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	if founduser.Email != "" {
		founduser.Email = user.Email
	}
	if founduser.PhoneNumber != "" {
		founduser.PhoneNumber = user.PhoneNumber
	}
	if founduser.Password != "" {
		founduser.Password = user.Password
	}
	founduser.IsLocked = true
	result := ari.DB.Save(&founduser)
	if result.Error != nil {
		log.Println("User Not Updated maybe the same email already present, AdminRepositoryImpl package")
		return nil, errors.New("user not updated")
	}
	return founduser, nil
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminRepositoryImpl{
		DB: db,
	}
}
