package integration_tests

import (
	"fmt"
	"fullstack/backend/internal/entity"
	handler "fullstack/backend/internal/handler/http"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func initTests(t *testing.T) {
	srv := handler.InitializeServer(testURL, testDBName, fakeSigningKey)

	handler.StartPolling(srv)

	clearDB(t)
	createTestData(t)
}

func createTestData(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(testDBName), &gorm.Config{})
	if err != nil {
		assert.Fail(t, "failed to connect database")
	}

	users := []entity.User{
		{
			ID: 1,
			UserRegister: entity.UserRegister{
				UserLogin: entity.UserLogin{
					Email:    "test1@gmail.com",
					Password: "1234",
				},
			},
		},
		{
			ID: 2,
			UserRegister: entity.UserRegister{
				UserLogin: entity.UserLogin{
					Email:    "test2@gmail.com",
					Password: "1234",
				},
			},
		},
	}
	items := []entity.Item{
		{SellerID: 1, Name: "IPhone 14", Description: "Text", ImageUrl: "iphone.jpg", Cost: 199.99},
		{SellerID: 2, Name: "IPhone 15", Description: "Text", ImageUrl: "iphone.jpg", Cost: 299.99},
		{SellerID: 2, Name: "IPhone 16", Description: "Text", ImageUrl: "iphone.jpg", Cost: 399.99},
	}
	orders := []entity.Order{
		{ID: 1, UserID: 1, ItemID: 1, Count: 2},
		{ID: 2, UserID: 2, ItemID: 2, Count: 3},
		{ID: 3, UserID: 2, ItemID: 3, Count: 4},
	}

	fmt.Println("Filling DB", users, items, orders)
	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
	for _, order := range orders {
		db.Create(&order)
	}
}

func clearDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(testDBName), &gorm.Config{})
	if err != nil {
		assert.Fail(t, "failed to connect database")
	}

	db.Unscoped().Where("1=1").Delete(&entity.Item{})
	db.Unscoped().Where("1=1").Delete(&entity.Order{})
	db.Unscoped().Where("1=1").Delete(&entity.User{})
}
