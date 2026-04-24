package models

import (
	"time"

	"gorm.io/gorm"
)

// User merepresentasikan data pengguna dalam database
type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(100);not null"`
	PhoneNumber string         `json:"phone_number" gorm:"type:varchar(20);not null;uniqueIndex"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete — disembunyikan dari JSON response
}

// UserResponse adalah DTO untuk response JSON (tanpa field internal seperti DeletedAt)
type UserResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToResponse mengkonversi User model ke UserResponse DTO
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

// ToResponseList mengkonversi slice User ke slice UserResponse
func ToUserResponseList(users []User) []UserResponse {
	responses := make([]UserResponse, len(users))
	for i, u := range users {
		responses[i] = u.ToResponse()
	}
	return responses
}

// CreateUserRequest adalah request body untuk membuat user baru
type CreateUserRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=15"`
}

// UpdateUserRequest adalah request body untuk mengupdate user
type UpdateUserRequest struct {
	Name        string `json:"name" validate:"omitempty,min=3,max=100"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,min=10,max=15"`
}
