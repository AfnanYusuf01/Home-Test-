package models

// PaketData merepresentasikan data paket internet
type PaketData struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"not null"`
	Price        float64 `json:"price" gorm:"not null"`
	Quota        float64 `json:"quota" gorm:"not null"`          // dalam GB
	ActivePeriod int     `json:"active_period" gorm:"not null"`   // dalam hari
}

// CreatePaketDataRequest adalah request body untuk membuat paket data baru
type CreatePaketDataRequest struct {
	Name         string  `json:"name" validate:"required"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	Quota        float64 `json:"quota" validate:"required,gt=0"`
	ActivePeriod int     `json:"active_period" validate:"required,gt=0"`
}

// UpdatePaketDataRequest adalah request body untuk mengupdate paket data
type UpdatePaketDataRequest struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Quota        float64 `json:"quota"`
	ActivePeriod int     `json:"active_period"`
}
