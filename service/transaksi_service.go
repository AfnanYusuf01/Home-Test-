package service

import (
	"github.com/AfnanYusuf01/take-home-test/helpers"
	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/repository"
)

// GetAllTransaksi mengambil semua transaksi beserta relasi User dan PaketData
func GetAllTransaksi() ([]models.Transaksi, error) {
	transaksi, err := repository.GetAllTransaksi()
	if err != nil {
		return nil, helpers.NewInternalError("Gagal mengambil data transaksi")
	}
	return transaksi, nil
}

// GetTransaksiByID mengambil transaksi berdasarkan ID beserta relasi
func GetTransaksiByID(id uint) (models.Transaksi, error) {
	transaksi, err := repository.GetTransaksiByID(id)
	if err != nil {
		return models.Transaksi{}, helpers.NewNotFoundError("Transaksi dengan ID tersebut tidak ditemukan")
	}
	return transaksi, nil
}

// CreateTransaksi membuat transaksi baru (user membeli paket data).
// Harga otomatis diambil dari paket data saat transaksi dibuat.
// Hanya user dan paket data yang aktif (belum di-soft-delete) yang bisa digunakan.
func CreateTransaksi(req models.CreateTransaksiRequest) (models.Transaksi, error) {
	if err := helpers.ValidateStruct(req); err != nil {
		return models.Transaksi{}, err
	}

	// Cek apakah user ada dan masih aktif (belum di-soft-delete)
	if _, err := repository.GetUserByID(req.UserID); err != nil {
		return models.Transaksi{}, helpers.NewNotFoundError("User dengan ID tersebut tidak ditemukan atau sudah tidak aktif")
	}

	// Cek apakah paket data ada dan masih aktif, sekaligus ambil harganya
	paketData, err := repository.GetPaketDataByID(req.PaketDataID)
	if err != nil {
		return models.Transaksi{}, helpers.NewNotFoundError("Paket data dengan ID tersebut tidak ditemukan atau sudah tidak aktif")
	}

	// Buat transaksi — harga otomatis diambil dari paket data saat ini
	transaksi := models.Transaksi{
		UserID:      req.UserID,
		PaketDataID: req.PaketDataID,
		Price:       paketData.Price,
	}

	if err := repository.CreateTransaksi(&transaksi); err != nil {
		return models.Transaksi{}, helpers.NewInternalError("Gagal membuat transaksi")
	}

	// Ambil transaksi lengkap dengan relasi User dan PaketData
	result, err := repository.GetTransaksiByID(transaksi.ID)
	if err != nil {
		return models.Transaksi{}, helpers.NewInternalError("Transaksi berhasil dibuat tetapi gagal mengambil detail")
	}

	return result, nil
}
