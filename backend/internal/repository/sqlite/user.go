package reposqlite

import (
	"errors"
	"fullstack/backend/internal/entity"
	"gorm.io/gorm"
	"log"
)

type UserSQLite struct {
	db *gorm.DB
}

func NewUserSQLite(db *gorm.DB) *UserSQLite {
	return &UserSQLite{db: db}
}

func (u UserSQLite) Create(user *entity.User) error {
	if result := u.db.Create(user); result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrDuplicatedKey):
			return entity.ErrUserExists
		default:
			return result.Error
		}
	} else {
		log.Println("Create result: ", result)
		return nil
	}
}

func (u UserSQLite) Update(user *entity.User) error {
	result := u.db.Model(user).Updates(user)
	if result.Error == nil {
		return nil
	}
	log.Println("Update result: ", result)
	return result.Error
}

func (u UserSQLite) Delete(id uint) error {
	result := u.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Delete result: ", result)
	return nil
}

func (u UserSQLite) Get(id uint) (*entity.User, error) {
	var user entity.User

	result := u.db.Where("id = ?", id).First(&user)
	if result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrUserNotFound
	}
	log.Println("Get result: ", user)
	return nil, result.Error
}

func (u UserSQLite) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	result := u.db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user, entity.ErrUserNotFound
	}
	log.Println("GetByEmail result: ", user)
	return &user, result.Error
}

func (u UserSQLite) GetAll() (*[]entity.User, error) {
	var users []entity.User

	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("GetAll result: ", users)
	return &users, nil
}
