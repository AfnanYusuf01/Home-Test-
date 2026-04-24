package models

import "time"

// User merepresentasikan data pengguna
type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"not null;uniqueIndex"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateUserRequest adalah request body untuk membuat user baru
type CreateUserRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

// UpdateUserRequest adalah request body untuk mengupdate user
type UpdateUserRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
