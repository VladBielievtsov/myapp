package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"my-app/db"
	"my-app/types"
	"strings"
	"time"
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
		Username: strings.ToLower(req.Username),
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

func (s *UserService) Login(req types.LoginUserRequest) (types.User, string, map[string]string) {
	if req.Username == "" || req.Password == "" {
		return types.User{}, "", map[string]string{"status": "fail", "message": "Username and password are required"}
	}

	var user types.User
	result := db.DB.First(&user, "username = ?", strings.ToLower(req.Username))
	if result.Error != nil {
		return types.User{}, "", map[string]string{"status": "fail", "message": "User not found"}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		if result.Error != nil {
			return types.User{}, "", map[string]string{"status": "fail", "message": "Invalid password"}
		}
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(120 * time.Minute).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte("secret"))
	if err != nil {
		return types.User{}, "", map[string]string{"status": "fail", "message": "generating JWT Token failed"}
	}

	return user, tokenString, nil
}
