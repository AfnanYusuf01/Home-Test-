package repository

import (
	"github.com/AfnanYusuf01/take-home-test/config"
	"github.com/AfnanYusuf01/take-home-test/models"
	"gorm.io/gorm"
)

// GetAllTransaksi mengambil semua data transaksi beserta relasi User dan PaketData.
// Menggunakan Unscoped pada Preload agar data User/PaketData yang sudah di-soft-delete
// tetap ditampilkan (karena transaksi adalah catatan historis).
func GetAllTransaksi() ([]models.Transaksi, error) {
	var transaksi []models.Transaksi
	result := config.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Unscoped() }).
		Preload("PaketData", func(db *gorm.DB) *gorm.DB { return db.Unscoped() }).
		Find(&transaksi)
	return transaksi, result.Error
}

// GetTransaksiByID mengambil data transaksi berdasarkan ID beserta relasi.
// Menggunakan Unscoped pada Preload untuk konsistensi dengan GetAllTransaksi.
func GetTransaksiByID(id uint) (models.Transaksi, error) {
	var transaksi models.Transaksi
	result := config.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Unscoped() }).
		Preload("PaketData", func(db *gorm.DB) *gorm.DB { return db.Unscoped() }).
		First(&transaksi, id)
	return transaksi, result.Error
}

// CreateTransaksi membuat transaksi baru di database
func CreateTransaksi(transaksi *models.Transaksi) error {
	return config.DB.Create(transaksi).Error
}
