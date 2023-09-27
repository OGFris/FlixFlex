package handlers

import (
	"github.com/OGFris/FlixFlex/internal/models"
	"github.com/OGFris/FlixFlex/internal/services"
	"github.com/OGFris/FlixFlex/internal/transport/requests"
	"github.com/OGFris/FlixFlex/internal/transport/responses"
	"github.com/OGFris/FlixFlex/pkg/errors"
	"github.com/OGFris/FlixFlex/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	var req requests.UserLogin

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.ErrorResponse{Message: errors.ErrDataFormat.Error()})
	}

	if err := req.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.ErrorResponse{Message: err.Error()})
	}

	user, err := handler.AuthService.Login(req.Username, req.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(errors.ErrorResponse{Message: err.Error()})
	}

	accessToken, refreshToken, err := handler.AuthService.GenerateJWT(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(responses.UserJWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (handler *AuthHandler) Register(c *fiber.Ctx) error {
	var req requests.UserRegister

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.ErrorResponse{Message: errors.ErrDataFormat.Error()})
	}

	if err := req.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.ErrorResponse{Message: err.Error()})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.ErrorResponse{Message: err.Error()})
	}

	user := &models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	}

	if err := handler.AuthService.Register(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.ErrorResponse{Message: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}
