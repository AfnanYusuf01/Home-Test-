package handler

import (
	"strconv"

	"github.com/AfnanYusuf01/take-home-test/helpers"
	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GetAllPaketData menampilkan semua paket data aktif
// @route GET /api/paket-data
func GetAllPaketData(c *fiber.Ctx) error {
	paketData, err := service.GetAllPaketData()
	if err != nil {
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil semua data paket data", models.ToPaketDataResponseList(paketData))
}

// GetPaketDataByID menampilkan paket data berdasarkan ID
// @route GET /api/paket-data/:id
func GetPaketDataByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID harus berupa angka yang valid")
	}

	paketData, err := service.GetPaketDataByID(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil data paket data", paketData.ToResponse())
}

// CreatePaketData membuat paket data baru
// @route POST /api/paket-data
func CreatePaketData(c *fiber.Ctx) error {
	var req models.CreatePaketDataRequest

	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Format request body tidak valid")
	}

	paketData, err := service.CreatePaketData(req)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidationErrorResponse(c, err)
		}
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "Paket data berhasil dibuat", paketData.ToResponse())
}

// UpdatePaketData mengupdate data paket data
// @route PUT /api/paket-data/:id
func UpdatePaketData(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID harus berupa angka yang valid")
	}

	var req models.UpdatePaketDataRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Format request body tidak valid")
	}

	paketData, err := service.UpdatePaketData(uint(id), req)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidationErrorResponse(c, err)
		}
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Paket data berhasil diupdate", paketData.ToResponse())
}

// DeletePaketData melakukan soft delete paket data
// @route DELETE /api/paket-data/:id
func DeletePaketData(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID harus berupa angka yang valid")
	}

	if err := service.DeletePaketData(uint(id)); err != nil {
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Paket data berhasil dihapus (soft delete)", nil)
}
