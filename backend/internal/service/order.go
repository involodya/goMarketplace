package service

import "fullstack/backend/internal/entity"

type OrderService struct {
	orderRepo entity.OrderRepository
}

func NewOrderService(orderRepo entity.OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (o OrderService) Create(order *entity.Order) error {
	err := o.orderRepo.Create(order)
	return err
}

func (o OrderService) GetBySeller(sellerId uint) (*[]entity.Order, error) {
	res, err := o.orderRepo.GetBySeller(sellerId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
