package service

import "fullstack/backend/internal/entity"

type ItemService struct {
	itemRepo  entity.ItemRepository
	usersRepo entity.UserRepository
}

func NewItemService(itemRepo entity.ItemRepository, usersRepo entity.UserRepository) *ItemService {
	return &ItemService{
		itemRepo:  itemRepo,
		usersRepo: usersRepo,
	}
}

func (i ItemService) Create(item *entity.Item) error {
	_, err := i.usersRepo.Get(item.SellerID)
	if err != nil {
		return entity.ErrUserNotFound
	}

	err = i.itemRepo.Create(item)
	return err
}

func (i ItemService) Update(item *entity.Item) error {
	_, err := i.usersRepo.Get(item.SellerID)
	if err != nil {
		return entity.ErrUserNotFound
	}

	err = i.itemRepo.Update(item)
	return err
}

func (i ItemService) Get(id uint) (*entity.Item, error) {
	res, err := i.itemRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i ItemService) GetBySeller(sellerId uint) (*[]entity.Item, error) {
	res, err := i.itemRepo.GetBySeller(sellerId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
