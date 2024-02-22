package reposqlite

import (
	"fullstack/backend/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB(dbUri string) (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Item{},
		&entity.Order{},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}
