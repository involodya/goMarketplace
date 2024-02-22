package entity

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	SellerID    uint
	Count       uint
	Name        string
	Description string
	ImageUrl    string
	Cost        float32
}

type ItemRepository interface {
	Create(*Item) error
	Update(*Item) error
	Delete(id uint) error

	Get(id uint) (*Item, error)
	GetBySeller(id uint) (*[]Item, error)
	GetAll() (*[]Item, error)
}

type ItemService interface {
	Create(*Item) error
	Update(*Item) error
	//Delete(id uint) error

	Get(id uint) (*Item, error)
	GetBySeller(id uint) (*[]Item, error)
}
