package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    int            `gorm:"primaryKey" json:"user_id"`
	Name      string         `json:"name" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	Update(user *User) (bool, error)
	FindByID(ID int) (*User, error)
	Delete(user *User) (bool, error)
}

type UserService interface {
	Create(user *User) (*User, error)
	Update(ID int, user *User) (bool, error)
	Delete(ID int) (bool, error)
}
