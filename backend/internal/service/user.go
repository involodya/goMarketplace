package service

import (
	"fullstack/backend/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo  entity.UserRepository
	usersRepo entity.UserRepository
}

func NewUserService(userRepo entity.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Get(id uint) (*entity.User, error) {
	user, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByEmail(email string) (*entity.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *UserService) GetAll() (*[]entity.User, error) {
	user, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func comparePasswordWithHash(password, oldPasswordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(oldPasswordHash), []byte(password))
}

func generatePasswordHash(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(res), err
}

func (s *UserService) Register(userReg *entity.UserRegister) error {
	if userReg.Email == "" {
		return entity.ErrInvalidEmail
	}

	if len(userReg.Password) < 8 {
		return entity.ErrInvalidPassword
	}

	pswdHash, err := generatePasswordHash(userReg.Password)
	if err != nil {
		return err
	}

	userReg.Password = pswdHash
	user := entity.User{UserRegister: *userReg}
	err = s.userRepo.Create(&user)
	return err
}

func (s *UserService) Login(user *entity.UserLogin) (uint, error) {
	userInDb, err := s.userRepo.GetByEmail(user.Email)
	if err != nil {
		return 0, err
	}

	err = comparePasswordWithHash(user.Password, userInDb.Password)
	if err != nil {
		return 0, err
	}

	return userInDb.ID, nil
}
