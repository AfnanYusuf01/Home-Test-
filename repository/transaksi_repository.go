package repository

import (
	"github.com/AfnanYusuf01/take-home-test/config"
	"github.com/AfnanYusuf01/take-home-test/models"
)

// GetAllTransaksi mengambil semua data transaksi beserta relasi User dan PaketData
func GetAllTransaksi() ([]models.Transaksi, error) {
	var transaksi []models.Transaksi
	result := config.DB.Preload("User").Preload("PaketData").Find(&transaksi)
	return transaksi, result.Error
}

// GetTransaksiByID mengambil data transaksi berdasarkan ID beserta relasi
func GetTransaksiByID(id uint) (models.Transaksi, error) {
	var transaksi models.Transaksi
	result := config.DB.Preload("User").Preload("PaketData").First(&transaksi, id)
	return transaksi, result.Error
}

// CreateTransaksi membuat transaksi baru di database
func CreateTransaksi(transaksi *models.Transaksi) error {
	return config.DB.Create(transaksi).Error
}
