package routes

import (
	"github.com/OGFris/FlixFlex/internal/handlers"
	"github.com/OGFris/FlixFlex/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"os"
)

func RegisterRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	// Authentication routes
	authRoutes := app.Group("/auth")
	authRoutes.Post("/login", authHandler.Login)
	authRoutes.Post("/register", authHandler.Register)

	// Regular user routes
	userRoutes := app.Group("/user")
	userRoutes.Use(middleware.UseJWT(os.Getenv("JWT_KEY")))

	// Other routes (public, etc.)

}
