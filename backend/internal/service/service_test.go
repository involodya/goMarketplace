package service_test

import (
	"fullstack/backend/internal/entity"
	"fullstack/backend/internal/service"
	mocks "fullstack/backend/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockUserId := uint(1)
	mockUser := entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Email:    "mock@gmail.com",
				Password: "123456",
			},
		},
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", mockUserId).Return(&mockUser, nil)

		s := service.NewUserService(mockUserRepo)
		user, err := s.Get(mockUserId)

		assert.NoError(t, err)
		assert.Equal(t, user, &mockUser)
		mockUserRepo.AssertExpectations(t)
	})
}
