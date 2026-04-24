package main

import (
	"log"
	"os"

	"github.com/AfnanYusuf01/take-home-test/config"
	"github.com/AfnanYusuf01/take-home-test/helpers"
	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variables sistem")
	}

	// Koneksi ke database
	config.ConnectDatabase()

	// Auto migrate semua model
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.PaketData{},
		&models.Transaksi{},
	)
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
	log.Println("Migrasi database berhasil")

	// Inisialisasi Fiber app dengan custom error handler
	app := fiber.New(fiber.Config{
		AppName: "Paket Data API v1.0",
		// Custom error handler untuk menangkap error yang tidak tertangani
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "Terjadi kesalahan internal server"

			// Cek apakah error dari Fiber (misal: 404 route not found)
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				message = e.Message
			}

			// Cek apakah error dari AppError custom kita
			if e, ok := err.(*helpers.AppError); ok {
				code = e.Code
				message = e.Message
			}

			return helpers.ErrorResponse(c, code, message)
		},
	})

	// Middleware
	app.Use(recover.New())  // Recovery dari panic
	app.Use(logger.New())   // HTTP request logger
	app.Use(cors.New())     // Cross-Origin Resource Sharing

	// Welcome route
	app.Get("/", func(c *fiber.Ctx) error {
		return helpers.SuccessResponse(c, fiber.StatusOK, "Selamat datang di Paket Data API", fiber.Map{
			"version": "1.0.0",
			"endpoints": fiber.Map{
				"users":      "/api/users",
				"paket_data": "/api/paket-data",
				"transaksi":  "/api/transaksi",
			},
		})
	})

	// Setup routes
	routes.SetupRoutes(app)

	// Tentukan port
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	// Jalankan server
	log.Printf("Server berjalan di http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
