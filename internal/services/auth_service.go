package services

import (
	"github.com/OGFris/FlixFlex/internal/models"
	"github.com/OGFris/FlixFlex/internal/repositories"
	"github.com/OGFris/FlixFlex/pkg/errors"
	"github.com/OGFris/FlixFlex/pkg/middleware"
	"github.com/OGFris/FlixFlex/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type AuthService interface {
	Login(username, password string) (*models.User, error)
	GenerateJWT(user *models.User) (string, string, error)
}

type UserAuthService struct {
	UserRepository repositories.UserRepository
}

func NewUserAuthService(userRepository repositories.UserRepository) *UserAuthService {
	return &UserAuthService{UserRepository: userRepository}
}

func (s *UserAuthService) Login(username, password string) (*models.User, error) {
	user, err := s.UserRepository.FindByUsername(username)
	if err != nil {

		return nil, err
	}

	if !utils.ComparePasswords(password, user.Password) {

		return nil, errors.ErrWrongPassword
	}

	return user, nil
}

func (s *UserAuthService) GenerateJWT(user *models.User) (string, string, error) {
	claims := middleware.UserClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		},
	}
	refreshClaims := middleware.UserRefreshClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {

		return "", "", errors.ErrJWTError
	}
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refresh.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {

		return "", "", errors.ErrJWTError
	}

	return tokenString, refreshString, nil
}
