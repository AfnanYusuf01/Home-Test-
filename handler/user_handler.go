package handler

import (
	"strconv"

	"github.com/AfnanYusuf01/take-home-test/helpers"
	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GetAllUsers menampilkan semua user aktif
// @route GET /api/users
func GetAllUsers(c *fiber.Ctx) error {
	users, err := service.GetAllUsers()
	if err != nil {
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil semua data user", models.ToUserResponseList(users))
}

// GetUserByID menampilkan user berdasarkan ID
// @route GET /api/users/:id
func GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID harus berupa angka yang valid")
	}

	user, err := service.GetUserByID(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil data user", user.ToResponse())
}

// CreateUser membuat user baru
// @route POST /api/users
func CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Format request body tidak valid")
	}

	user, err := service.CreateUser(req)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidationErrorResponse(c, err)
		}
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "User berhasil dibuat", user.ToResponse())
}

// UpdateUser mengupdate data user
// @route PUT /api/users/:id
func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID harus berupa angka yang valid")
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Format request body tidak valid")
	}

	user, err := service.UpdateUser(uint(id), req)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidationErrorResponse(c, err)
		}
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "User berhasil diupdate", user.ToResponse())
}

// DeleteUser melakukan soft delete user
// @route DELETE /api/users/:id
func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID harus berupa angka yang valid")
	}

	if err := service.DeleteUser(uint(id)); err != nil {
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "User berhasil dihapus (soft delete)", nil)
}
