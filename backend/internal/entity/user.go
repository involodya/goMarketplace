package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserRegister
}

type UserLogin struct {
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null" json:"Password,omitempty"`
}

type UserRegister struct {
	UserLogin

	//FirstName string
	//LastName  string
}

type UserRepository interface {
	Create(*User) error
	Update(*User) error
	Delete(id uint) error

	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() (*[]User, error)
}

type UserService interface {
	Register(userReg *UserRegister) error
	Login(userLogin *UserLogin) (uint, error)

	Get(id uint) (*User, error)
}
