package handler

import (
	"strconv"

	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/service"
	"github.com/gofiber/fiber/v2"
)

// GetAllTransaksi menampilkan semua transaksi
// @route GET /api/transaksi
func GetAllTransaksi(c *fiber.Ctx) error {
	transaksi, err := service.GetAllTransaksi()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data transaksi",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua data transaksi",
		"data":    transaksi,
	})
}

// GetTransaksiByID menampilkan transaksi berdasarkan ID
// @route GET /api/transaksi/:id
func GetTransaksiByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
		})
	}

	transaksi, err := service.GetTransaksiByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Transaksi tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil data transaksi",
		"data":    transaksi,
	})
}

// CreateTransaksi membuat transaksi baru (user membeli paket data)
// @route POST /api/transaksi
func CreateTransaksi(c *fiber.Ctx) error {
	var req models.CreateTransaksiRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	transaksi, err := service.CreateTransaksi(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Transaksi berhasil dibuat",
		"data":    transaksi,
	})
}
