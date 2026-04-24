package handler

import (
	"strconv"

	"github.com/AfnanYusuf01/take-home-test/helpers"
	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GetAllTransaksi menampilkan semua transaksi beserta relasi
// @route GET /api/transaksi
func GetAllTransaksi(c *fiber.Ctx) error {
	transaksi, err := service.GetAllTransaksi()
	if err != nil {
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil semua data transaksi", models.ToTransaksiResponseList(transaksi))
}

// GetTransaksiByID menampilkan transaksi berdasarkan ID
// @route GET /api/transaksi/:id
func GetTransaksiByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID harus berupa angka yang valid")
	}

	transaksi, err := service.GetTransaksiByID(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil data transaksi", transaksi.ToResponse())
}

// CreateTransaksi membuat transaksi baru (user membeli paket data)
// @route POST /api/transaksi
func CreateTransaksi(c *fiber.Ctx) error {
	var req models.CreateTransaksiRequest

	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Format request body tidak valid")
	}

	transaksi, err := service.CreateTransaksi(req)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidationErrorResponse(c, err)
		}
		return helpers.ErrorResponse(c, helpers.GetStatusCode(err), err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "Transaksi berhasil dibuat", transaksi.ToResponse())
}
