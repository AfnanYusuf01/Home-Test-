package service

import (
	"errors"

	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/repository"
)

// GetAllPaketData mengambil semua paket data
func GetAllPaketData() ([]models.PaketData, error) {
	return repository.GetAllPaketData()
}

// GetPaketDataByID mengambil paket data berdasarkan ID
func GetPaketDataByID(id uint) (models.PaketData, error) {
	return repository.GetPaketDataByID(id)
}

// CreatePaketData membuat paket data baru dengan validasi
func CreatePaketData(req models.CreatePaketDataRequest) (models.PaketData, error) {
	if req.Name == "" {
		return models.PaketData{}, errors.New("nama paket tidak boleh kosong")
	}
	if req.Price <= 0 {
		return models.PaketData{}, errors.New("harga harus lebih dari 0")
	}
	if req.Quota <= 0 {
		return models.PaketData{}, errors.New("kuota harus lebih dari 0")
	}
	if req.ActivePeriod <= 0 {
		return models.PaketData{}, errors.New("masa aktif harus lebih dari 0")
	}

	paketData := models.PaketData{
		Name:         req.Name,
		Price:        req.Price,
		Quota:        req.Quota,
		ActivePeriod: req.ActivePeriod,
	}

	err := repository.CreatePaketData(&paketData)
	return paketData, err
}

// UpdatePaketData mengupdate data paket data dengan validasi
func UpdatePaketData(id uint, req models.UpdatePaketDataRequest) (models.PaketData, error) {
	paketData, err := repository.GetPaketDataByID(id)
	if err != nil {
		return models.PaketData{}, errors.New("paket data tidak ditemukan")
	}

	if req.Name != "" {
		paketData.Name = req.Name
	}
	if req.Price > 0 {
		paketData.Price = req.Price
	}
	if req.Quota > 0 {
		paketData.Quota = req.Quota
	}
	if req.ActivePeriod > 0 {
		paketData.ActivePeriod = req.ActivePeriod
	}

	err = repository.UpdatePaketData(&paketData)
	return paketData, err
}

// DeletePaketData menghapus paket data berdasarkan ID
func DeletePaketData(id uint) error {
	// Cek apakah paket data ada
	_, err := repository.GetPaketDataByID(id)
	if err != nil {
		return errors.New("paket data tidak ditemukan")
	}

	return repository.DeletePaketData(id)
}
