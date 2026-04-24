package service

import (
	"strings"

	"github.com/AfnanYusuf01/take-home-test/helpers"
	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/repository"
)

// GetAllPaketData mengambil semua paket data aktif (yang belum di-soft-delete)
func GetAllPaketData() ([]models.PaketData, error) {
	paketData, err := repository.GetAllPaketData()
	if err != nil {
		return nil, helpers.NewInternalError("Gagal mengambil data paket data")
	}
	return paketData, nil
}

// GetPaketDataByID mengambil paket data aktif berdasarkan ID
func GetPaketDataByID(id uint) (models.PaketData, error) {
	paketData, err := repository.GetPaketDataByID(id)
	if err != nil {
		return models.PaketData{}, helpers.NewNotFoundError("Paket data dengan ID tersebut tidak ditemukan")
	}
	return paketData, nil
}

// CreatePaketData membuat paket data baru dengan validasi
func CreatePaketData(req models.CreatePaketDataRequest) (models.PaketData, error) {
	if err := helpers.ValidateStruct(req); err != nil {
		return models.PaketData{}, err
	}

	paketData := models.PaketData{
		Name:         strings.TrimSpace(req.Name),
		Price:        req.Price,
		Quota:        req.Quota,
		ActivePeriod: req.ActivePeriod,
	}

	if err := repository.CreatePaketData(&paketData); err != nil {
		return models.PaketData{}, helpers.NewInternalError("Gagal membuat paket data")
	}

	return paketData, nil
}

// UpdatePaketData mengupdate data paket data dengan validasi partial update
func UpdatePaketData(id uint, req models.UpdatePaketDataRequest) (models.PaketData, error) {
	if err := helpers.ValidateStruct(req); err != nil {
		return models.PaketData{}, err
	}

	paketData, err := repository.GetPaketDataByID(id)
	if err != nil {
		return models.PaketData{}, helpers.NewNotFoundError("Paket data dengan ID tersebut tidak ditemukan")
	}

	if req.Name != "" {
		paketData.Name = strings.TrimSpace(req.Name)
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

	if err := repository.UpdatePaketData(&paketData); err != nil {
		return models.PaketData{}, helpers.NewInternalError("Gagal mengupdate paket data")
	}

	return paketData, nil
}

// DeletePaketData melakukan soft delete pada paket data berdasarkan ID.
// Data transaksi historis tetap tersimpan untuk keperluan audit.
func DeletePaketData(id uint) error {
	if _, err := repository.GetPaketDataByID(id); err != nil {
		return helpers.NewNotFoundError("Paket data dengan ID tersebut tidak ditemukan")
	}

	if err := repository.DeletePaketData(id); err != nil {
		return helpers.NewInternalError("Gagal menghapus paket data")
	}

	return nil
}
