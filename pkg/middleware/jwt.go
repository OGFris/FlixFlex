package middleware

import (
	"fmt"
	"github.com/OGFris/FlixFlex/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

type (
	JWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}
	Skipper      func(c *fiber.Ctx) bool
	jwtExtractor func(*fiber.Ctx) (string, error)
)

func UseJWT(key interface{}) fiber.Handler {
	c := JWTConfig{}
	c.SigningKey = key
	return JWTWithConfig(c)
}

func JWTWithConfig(config JWTConfig) fiber.Handler {
	extractor := jwtFromHeader("Authorization", "Bearer")
	return func(c *fiber.Ctx) error {
		auth, err := extractor(c)
		if err != nil {
			if config.Skipper != nil && config.Skipper(c) {
				return c.Next()
			}
			return c.Status(http.StatusUnauthorized).JSON(errors.ErrorResponse{Message: errors.ErrJWTMissing.Error()})
		}

		claims, err := validateToken(auth, config.SigningKey.(string))
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(errors.ErrorResponse{errors.ErrJWTInvalid.Error()})
		}

		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}

func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c *fiber.Ctx) (string, error) {
		auth := c.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", errors.ErrJWTMissing
	}
}

func validateToken(tokenString string, signingKey string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.ErrJWTInvalidClaims
	}

	return claims, nil
}
