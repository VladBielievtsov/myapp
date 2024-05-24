package services

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return types.User{}, map[string]string{"status": "fail", "message": "Error during password hashing"}
	}

	id := uuid.New()
	user := types.User{
		ID:       &id,
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return types.User{}, map[string]string{"status": "fail", "message": "Could not create user"}
	}

	return user, nil
}

func (s *UserService) GetUserByID(id string) (types.User, map[string]string) {
	var user types.User

	result := db.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return types.User{}, map[string]string{"status": "fail", "message": "User not found"}
	}

	return user, nil
}
