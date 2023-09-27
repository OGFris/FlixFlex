package main

import (
	"github.com/OGFris/FlixFlex/internal/config"
	"github.com/OGFris/FlixFlex/internal/handlers"
	"github.com/OGFris/FlixFlex/internal/models"
	"github.com/OGFris/FlixFlex/internal/repositories"
	"github.com/OGFris/FlixFlex/internal/routes"
	"github.com/OGFris/FlixFlex/internal/services"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Movie{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate the database: %v", err)
	}

	// Register repositories
	userRepo := repositories.NewUserRepository(db)

	// Register services
	userAuthService := services.NewUserAuthService(userRepo)

	// Register handlers
	userAuthHandler := handlers.NewAuthHandler(userAuthService)

	// Create a new Fiber app
	app := fiber.New()

	// Register routes
	routes.RegisterRoutes(app, userAuthHandler)

	// Start the server
	log.Println("Server is running on http://localhost:" + cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
