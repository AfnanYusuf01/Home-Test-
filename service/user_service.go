package service

import (
	"errors"

	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/repository"
)

// GetAllUsers mengambil semua user
func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

// GetUserByID mengambil user berdasarkan ID
func GetUserByID(id uint) (models.User, error) {
	return repository.GetUserByID(id)
}

// CreateUser membuat user baru dengan validasi
func CreateUser(req models.CreateUserRequest) (models.User, error) {
	if req.Name == "" {
		return models.User{}, errors.New("nama tidak boleh kosong")
	}
	if req.PhoneNumber == "" {
		return models.User{}, errors.New("nomor telepon tidak boleh kosong")
	}

	user := models.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}

	err := repository.CreateUser(&user)
	return user, err
}

// UpdateUser mengupdate data user dengan validasi
func UpdateUser(id uint, req models.UpdateUserRequest) (models.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return models.User{}, errors.New("user tidak ditemukan")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}

	err = repository.UpdateUser(&user)
	return user, err
}

// DeleteUser menghapus user berdasarkan ID
func DeleteUser(id uint) error {
	// Cek apakah user ada
	_, err := repository.GetUserByID(id)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	return repository.DeleteUser(id)
}
