package main

import (
	"log"
	"os"

	"github.com/AfnanYusuf01/take-home-test/config"
	"github.com/AfnanYusuf01/take-home-test/models"
	"github.com/AfnanYusuf01/take-home-test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  File .env tidak ditemukan, menggunakan environment variables sistem")
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
	log.Println("✅ Migrasi database berhasil")

	// Inisialisasi Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Paket Data API v1.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Welcome route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "🚀 Selamat datang di Paket Data API",
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
	log.Printf("🚀 Server berjalan di http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
