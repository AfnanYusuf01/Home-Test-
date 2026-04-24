package models

import "time"

// Transaksi merepresentasikan transaksi pembelian paket data oleh user
type Transaksi struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null;index"`
	PaketDataID uint      `json:"paket_data_id" gorm:"not null;index"`
	Price       float64   `json:"price" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`

	// Relasi
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	PaketData PaketData `json:"paket_data" gorm:"foreignKey:PaketDataID"`
}

// CreateTransaksiRequest adalah request body untuk membuat transaksi baru
type CreateTransaksiRequest struct {
	UserID      uint `json:"user_id" validate:"required"`
	PaketDataID uint `json:"paket_data_id" validate:"required"`
}
