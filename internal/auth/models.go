package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null" validate:"required,email"`
	Password string `gorm:"not null" validate:"required,min=8"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserDB interface abstracts the database operations required by the auth handlers.
type UserDB interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}

// RealUserDB implements UserDB using Gorm
type RealUserDB struct {
	DB *gorm.DB
}

func (db *RealUserDB) CreateUser(user *User) error {
	return db.DB.Create(user).Error
}

func (db *RealUserDB) GetUserByEmail(email string) (*User, error) {
	var user User
	result := db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
