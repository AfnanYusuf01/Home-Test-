package routes

import (
	"github.com/AfnanYusuf01/take-home-test/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes mengatur semua routing endpoint API
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	users := api.Group("/users")
	users.Get("/", handler.GetAllUsers)
	users.Get("/:id", handler.GetUserByID)
	users.Post("/", handler.CreateUser)
	users.Put("/:id", handler.UpdateUser)
	users.Delete("/:id", handler.DeleteUser)

	// Paket Data routes
	paketData := api.Group("/paket-data")
	paketData.Get("/", handler.GetAllPaketData)
	paketData.Get("/:id", handler.GetPaketDataByID)
	paketData.Post("/", handler.CreatePaketData)
	paketData.Put("/:id", handler.UpdatePaketData)
	paketData.Delete("/:id", handler.DeletePaketData)

	// Transaksi routes
	transaksi := api.Group("/transaksi")
	transaksi.Get("/", handler.GetAllTransaksi)
	transaksi.Get("/:id", handler.GetTransaksiByID)
	transaksi.Post("/", handler.CreateTransaksi)
}
