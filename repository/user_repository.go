package repository

import (
	"github.com/AfnanYusuf01/take-home-test/config"
	"github.com/AfnanYusuf01/take-home-test/models"
)

// GetAllUsers mengambil semua data user dari database
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

// GetUserByID mengambil data user berdasarkan ID
func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

// CreateUser membuat user baru di database
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// UpdateUser mengupdate data user di database
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

// DeleteUser menghapus user dari database berdasarkan ID
func DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
