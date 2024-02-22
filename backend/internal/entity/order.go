package entity

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID uint
	ItemID uint
	//Status string
	Count uint
}

type OrderRepository interface {
	Create(*Order) error
	Update(*Order) error
	Delete(id uint) error

	Get(id uint) (*Order, error)
	GetBySeller(sellerId uint) (*[]Order, error)
	GetAll() (*[]Order, error)
}

type OrderService interface {
	Create(*Order) error
	//Update(*Order) error
	//Delete(id uint) error

	GetBySeller(sellerId uint) (*[]Order, error)
}
