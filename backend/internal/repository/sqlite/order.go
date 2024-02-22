package reposqlite

import (
	"errors"
	"fullstack/backend/internal/entity"
	"gorm.io/gorm"
	"log"
)

type OrderSQLite struct {
	db *gorm.DB
}

func NewOrderSQLite(db *gorm.DB) *OrderSQLite {
	return &OrderSQLite{db: db}
}

func (u OrderSQLite) Create(order *entity.Order) error {
	result := u.db.Create(order)
	if result.Error == nil {
		return nil
	}
	log.Println("Create result: ", result)
	return result.Error
}

func (u OrderSQLite) Update(order *entity.Order) error {
	result := u.db.Model(order).Updates(order)
	if result.Error == nil {
		return nil
	}
	log.Println("Update result: ", result)
	return result.Error
}

func (u OrderSQLite) Delete(id uint) error {
	result := u.db.Delete(&entity.Order{}, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Delete result: ", result)
	return nil
}

func (u OrderSQLite) Get(id uint) (*entity.Order, error) {
	var order entity.Order

	result := u.db.Where("id = ?", id).First(&order)
	if result.Error == nil {
		return &order, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrOrderNotFound
	}
	log.Println("Get result: ", order)
	return nil, result.Error
}

func (u OrderSQLite) GetAll() (*[]entity.Order, error) {
	var orders []entity.Order

	result := u.db.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("GetAll result: ", orders)
	return &orders, nil
}

func (u OrderSQLite) GetBySeller(sellerId uint) (*[]entity.Order, error) {
	var orders []entity.Order

	result := u.db.Joins("JOIN items ON items.id = orders.item_id").Where("items.seller_id = ? AND orders.deleted_at IS NULL", sellerId).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("GetBySeller result: ", orders)
	return &orders, nil
}
