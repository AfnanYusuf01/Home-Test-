package repository

import (
	"github.com/AfnanYusuf01/take-home-test/config"
	"github.com/AfnanYusuf01/take-home-test/models"
)

// GetAllPaketData mengambil semua data paket data dari database
func GetAllPaketData() ([]models.PaketData, error) {
	var paketData []models.PaketData
	result := config.DB.Find(&paketData)
	return paketData, result.Error
}

// GetPaketDataByID mengambil data paket data berdasarkan ID
func GetPaketDataByID(id uint) (models.PaketData, error) {
	var paketData models.PaketData
	result := config.DB.First(&paketData, id)
	return paketData, result.Error
}

// CreatePaketData membuat paket data baru di database
func CreatePaketData(paketData *models.PaketData) error {
	return config.DB.Create(paketData).Error
}

// UpdatePaketData mengupdate data paket data di database
func UpdatePaketData(paketData *models.PaketData) error {
	return config.DB.Save(paketData).Error
}

// DeletePaketData menghapus paket data dari database berdasarkan ID
func DeletePaketData(id uint) error {
	return config.DB.Delete(&models.PaketData{}, id).Error
}
