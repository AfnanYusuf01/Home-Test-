package models

import "time"

// Transaksi merepresentasikan transaksi pembelian paket data oleh user
type Transaksi struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null;index"`
	PaketDataID uint      `json:"paket_data_id" gorm:"not null;index"`
	Price       float64   `json:"price" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`

	// Relasi — Tidak menggunakan CASCADE karena User & PaketData menggunakan soft delete.
	// Data transaksi tetap utuh sebagai catatan historis meskipun User/PaketData di-soft-delete.
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	PaketData PaketData `json:"paket_data" gorm:"foreignKey:PaketDataID"`
}

// TransaksiResponse adalah DTO untuk response JSON
type TransaksiResponse struct {
	ID          uint              `json:"id"`
	UserID      uint              `json:"user_id"`
	PaketDataID uint              `json:"paket_data_id"`
	Price       float64           `json:"price"`
	CreatedAt   time.Time         `json:"created_at"`
	User        UserResponse      `json:"user"`
	PaketData   PaketDataResponse `json:"paket_data"`
}

// ToResponse mengkonversi Transaksi model ke TransaksiResponse DTO
func (t *Transaksi) ToResponse() TransaksiResponse {
	return TransaksiResponse{
		ID:          t.ID,
		UserID:      t.UserID,
		PaketDataID: t.PaketDataID,
		Price:       t.Price,
		CreatedAt:   t.CreatedAt,
		User:        t.User.ToResponse(),
		PaketData:   t.PaketData.ToResponse(),
	}
}

// ToTransaksiResponseList mengkonversi slice Transaksi ke slice TransaksiResponse
func ToTransaksiResponseList(items []Transaksi) []TransaksiResponse {
	responses := make([]TransaksiResponse, len(items))
	for i, t := range items {
		responses[i] = t.ToResponse()
	}
	return responses
}

// CreateTransaksiRequest adalah request body untuk membuat transaksi baru
type CreateTransaksiRequest struct {
	UserID      uint `json:"user_id" validate:"required"`
	PaketDataID uint `json:"paket_data_id" validate:"required"`
}
