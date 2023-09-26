package errors

import "errors"

var (
	ErrUserNotFound     = errors.New("user.not.found")
	ErrDatabase         = errors.New("unknown.database.error")
	ErrWrongPassword    = errors.New("wrong.password")
	ErrJWTMissing       = errors.New("missing.jwt")
	ErrJWTInvalid       = errors.New("invalid.jwt")
	ErrJWTInvalidClaims = errors.New("invalid.jwt.claims")
	ErrJWTError         = errors.New("error.creating.jwt")
	ErrDataFormat       = errors.New("data.format.incorrect")
	ErrMissingData      = errors.New("missing.data")
	ErrUserDisabled     = errors.New("account.disabled")
	ErrNoAccess         = errors.New("no.access")
)

type ErrorResponse struct {
	Message string `json:"message"`
}
