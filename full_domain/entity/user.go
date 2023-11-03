package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `json:"email" gorm:"primary key"`
	Password    string `json:"password" gorm:"not null"`
	IsLocked    bool   `json:"is_account_locked" gorm:"default: false"`
	Role        string `json:"role" gorm:"default: 'user'"`
	PhoneNumber string `json:"phonenumber"`
}
