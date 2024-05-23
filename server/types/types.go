package types

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        *uuid.UUID `gorm:"type:uuid;not null;primary_key" json:"id,omitempty"`
	Username  string     `gorm:"varchar(255);unique;not null" json:"username,omitempty"`
	Password  string     `gorm:"varchar(255);not null" json:"password"`
	CreatedAt *time.Time `gorm:"not null;default:now()" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"not null;default:now()" json:"updatedAt"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FilterUser(user *User) UserResponse {
	return UserResponse{
		ID:        *user.ID,
		Username:  user.Username,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}
}

func FilterUsers(users []User) []UserResponse {
	var filteredUsers []UserResponse

	for _, user := range users {
		filteredUsers = append(filteredUsers, FilterUser(&user))
	}

	return filteredUsers
}
