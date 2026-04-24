package service

import (
	"errors"

	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/repository"
)

// GetAllTransaksi mengambil semua transaksi
func GetAllTransaksi() ([]models.Transaksi, error) {
	return repository.GetAllTransaksi()
}

// GetTransaksiByID mengambil transaksi berdasarkan ID
func GetTransaksiByID(id uint) (models.Transaksi, error) {
	return repository.GetTransaksiByID(id)
}

// CreateTransaksi membuat transaksi baru dengan validasi
func CreateTransaksi(req models.CreateTransaksiRequest) (models.Transaksi, error) {
	if req.UserID == 0 {
		return models.Transaksi{}, errors.New("user_id tidak boleh kosong")
	}
	if req.PaketDataID == 0 {
		return models.Transaksi{}, errors.New("paket_data_id tidak boleh kosong")
	}

	// Cek apakah user ada
	_, err := repository.GetUserByID(req.UserID)
	if err != nil {
		return models.Transaksi{}, errors.New("user tidak ditemukan")
	}

	// Cek apakah paket data ada dan ambil harganya
	paketData, err := repository.GetPaketDataByID(req.PaketDataID)
	if err != nil {
		return models.Transaksi{}, errors.New("paket data tidak ditemukan")
	}

	// Buat transaksi dengan harga saat ini dari paket data
	transaksi := models.Transaksi{
		UserID:      req.UserID,
		PaketDataID: req.PaketDataID,
		Price:       paketData.Price,
	}

	err = repository.CreateTransaksi(&transaksi)
	if err != nil {
		return models.Transaksi{}, err
	}

	// Ambil transaksi lengkap dengan relasi
	return repository.GetTransaksiByID(transaksi.ID)
}
