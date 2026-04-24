package service

import (
	"strings"

	"github.com/AfnanYusuf01/take-home-test/helpers"
	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/repository"
)

// GetAllUsers mengambil semua user aktif (yang belum di-soft-delete)
func GetAllUsers() ([]models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, helpers.NewInternalError("Gagal mengambil data user")
	}
	return users, nil
}

// GetUserByID mengambil user aktif berdasarkan ID
func GetUserByID(id uint) (models.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return models.User{}, helpers.NewNotFoundError("User dengan ID tersebut tidak ditemukan")
	}
	return user, nil
}

// CreateUser membuat user baru dengan validasi
func CreateUser(req models.CreateUserRequest) (models.User, error) {
	if err := helpers.ValidateStruct(req); err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:        strings.TrimSpace(req.Name),
		PhoneNumber: strings.TrimSpace(req.PhoneNumber),
	}

	if err := repository.CreateUser(&user); err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "23505") {
			return models.User{}, helpers.NewConflictError("Nomor telepon sudah digunakan oleh user lain")
		}
		return models.User{}, helpers.NewInternalError("Gagal membuat user")
	}

	return user, nil
}

// UpdateUser mengupdate data user dengan validasi partial update
func UpdateUser(id uint, req models.UpdateUserRequest) (models.User, error) {
	if err := helpers.ValidateStruct(req); err != nil {
		return models.User{}, err
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		return models.User{}, helpers.NewNotFoundError("User dengan ID tersebut tidak ditemukan")
	}

	if req.Name != "" {
		user.Name = strings.TrimSpace(req.Name)
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = strings.TrimSpace(req.PhoneNumber)
	}

	if err := repository.UpdateUser(&user); err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "23505") {
			return models.User{}, helpers.NewConflictError("Nomor telepon sudah digunakan oleh user lain")
		}
		return models.User{}, helpers.NewInternalError("Gagal mengupdate user")
	}

	return user, nil
}

// DeleteUser melakukan soft delete pada user berdasarkan ID.
// Data transaksi historis tetap tersimpan untuk keperluan audit.
func DeleteUser(id uint) error {
	if _, err := repository.GetUserByID(id); err != nil {
		return helpers.NewNotFoundError("User dengan ID tersebut tidak ditemukan")
	}

	if err := repository.DeleteUser(id); err != nil {
		return helpers.NewInternalError("Gagal menghapus user")
	}

	return nil
}
