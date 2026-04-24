package models

import (
	"time"

	"gorm.io/gorm"
)

// PaketData merepresentasikan data paket internet dalam database
type PaketData struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"type:varchar(150);not null"`
	Price        float64        `json:"price" gorm:"not null"`
	Quota        float64        `json:"quota" gorm:"not null"`        // dalam GB
	ActivePeriod int            `json:"active_period" gorm:"not null"` // dalam hari
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete — disembunyikan dari JSON response
}

// PaketDataResponse adalah DTO untuk response JSON
type PaketDataResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Quota        float64   `json:"quota"`
	ActivePeriod int       `json:"active_period"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ToResponse mengkonversi PaketData model ke PaketDataResponse DTO
func (p *PaketData) ToResponse() PaketDataResponse {
	return PaketDataResponse{
		ID:           p.ID,
		Name:         p.Name,
		Price:        p.Price,
		Quota:        p.Quota,
		ActivePeriod: p.ActivePeriod,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

// ToPaketDataResponseList mengkonversi slice PaketData ke slice PaketDataResponse
func ToPaketDataResponseList(items []PaketData) []PaketDataResponse {
	responses := make([]PaketDataResponse, len(items))
	for i, p := range items {
		responses[i] = p.ToResponse()
	}
	return responses
}

// CreatePaketDataRequest adalah request body untuk membuat paket data baru
type CreatePaketDataRequest struct {
	Name         string  `json:"name" validate:"required,min=3,max=150"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	Quota        float64 `json:"quota" validate:"required,gt=0"`
	ActivePeriod int     `json:"active_period" validate:"required,gt=0"`
}

// UpdatePaketDataRequest adalah request body untuk mengupdate paket data
type UpdatePaketDataRequest struct {
	Name         string  `json:"name" validate:"omitempty,min=3,max=150"`
	Price        float64 `json:"price" validate:"omitempty,gt=0"`
	Quota        float64 `json:"quota" validate:"omitempty,gt=0"`
	ActivePeriod int     `json:"active_period" validate:"omitempty,gt=0"`
}
