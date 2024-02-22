package service

import (
	"fullstack/backend/internal/entity"
	"fullstack/backend/internal/repository"
)

type Service struct {
	User  entity.UserService
	Item  entity.ItemService
	Order entity.OrderService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:  NewUserService(repo.User),
		Item:  NewItemService(repo.Item, repo.User),
		Order: NewOrderService(repo.Order),
	}
}
