package repository

import (
	"fullstack/backend/internal/entity"
	reposqlite "fullstack/backend/internal/repository/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	User  entity.UserRepository
	Item  entity.ItemRepository
	Order entity.OrderRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:  reposqlite.NewUserSQLite(db),
		Item:  reposqlite.NewItemSQLite(db),
		Order: reposqlite.NewOrderSQLite(db),
	}
}
