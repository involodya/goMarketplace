package reposqlite

import (
	"errors"
	"fullstack/backend/internal/entity"
	"gorm.io/gorm"
	"log"
)

type ItemSQLite struct {
	db *gorm.DB
}

func NewItemSQLite(db *gorm.DB) *ItemSQLite {
	return &ItemSQLite{db: db}
}

func (u ItemSQLite) Create(item *entity.Item) error {
	result := u.db.Create(item)
	if result.Error == nil {
		return nil
	}
	log.Println("Create result: ", result)
	return result.Error
}

func (u ItemSQLite) Update(item *entity.Item) error {
	result := u.db.Model(item).Where("id = ?", item.ID).Updates(item)
	if result.Error == nil {
		return nil
	}
	log.Println("Update result: ", result)
	return result.Error
}

func (u ItemSQLite) Delete(id uint) error {
	result := u.db.Delete(&entity.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Delete result: ", result)
	return nil
}

func (u ItemSQLite) Get(id uint) (*entity.Item, error) {
	var item entity.Item

	result := u.db.Where("id = ?", id).First(&item)
	if result.Error == nil {
		return &item, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrItemNotFound
	}
	log.Println("Get result: ", item)
	return nil, result.Error
}

func (u ItemSQLite) GetBySeller(sellerId uint) (*[]entity.Item, error) {
	var items []entity.Item

	result := u.db.Where("seller_id = ?", sellerId).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("GetBySeller result: ", items)
	return &items, nil
}

func (u ItemSQLite) GetAll() (*[]entity.Item, error) {
	var items []entity.Item

	result := u.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("GetAll result: ", items)
	return &items, nil
}
