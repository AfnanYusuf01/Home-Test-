package handler

import (
	"strconv"

	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/service"
	"github.com/gofiber/fiber/v2"
)

// GetAllPaketData menampilkan semua paket data
// @route GET /api/paket-data
func GetAllPaketData(c *fiber.Ctx) error {
	paketData, err := service.GetAllPaketData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data paket data",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua data paket data",
		"data":    paketData,
	})
}

// GetPaketDataByID menampilkan paket data berdasarkan ID
// @route GET /api/paket-data/:id
func GetPaketDataByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
		})
	}

	paketData, err := service.GetPaketDataByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Paket data tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil data paket data",
		"data":    paketData,
	})
}

// CreatePaketData membuat paket data baru
// @route POST /api/paket-data
func CreatePaketData(c *fiber.Ctx) error {
	var req models.CreatePaketDataRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	paketData, err := service.CreatePaketData(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Paket data berhasil dibuat",
		"data":    paketData,
	})
}

// UpdatePaketData mengupdate data paket data
// @route PUT /api/paket-data/:id
func UpdatePaketData(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
		})
	}

	var req models.UpdatePaketDataRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Format request tidak valid",
			"error":   err.Error(),
		})
	}

	paketData, err := service.UpdatePaketData(uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Paket data berhasil diupdate",
		"data":    paketData,
	})
}

// DeletePaketData menghapus paket data
// @route DELETE /api/paket-data/:id
func DeletePaketData(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
		})
	}

	err = service.DeletePaketData(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Paket data berhasil dihapus",
	})
}
