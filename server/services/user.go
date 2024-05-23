package services

import (
	"github.com/google/uuid"
	"my-app/db"
	"my-app/types"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUsers() ([]types.User, map[string]string) {
	var users []types.User

	result := db.DB.Find(&users)
	if result.Error != nil {
		return nil, map[string]string{"status": "fail", "message": "Users not found"}
	}

	return users, nil
}

func (s *UserService) CreateUser(req types.CreateUserRequest) (types.User, map[string]string) {
	if req.Username == "" || req.Password == "" {
		return types.User{}, map[string]string{"status": "fail", "message": "Username and password are required"}
	}
	id := uuid.New()
	user := types.User{
		ID:       &id,
		Username: req.Username,
		Password: req.Password,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return types.User{}, map[string]string{"status": "fail", "message": "Could not create user"}
	}

	return user, nil
}
