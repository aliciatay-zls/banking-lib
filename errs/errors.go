package errs

import "net/http"

const MessageExpiredAccessToken = "expired access token"
const MessageInvalidAccessToken = "invalid access token"
const MessageRefreshToken = "expired or invalid refresh token"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewAuthenticationErrorDueToExpiredAccessToken() *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: MessageExpiredAccessToken,
	}
}

func NewAuthenticationErrorDueToInvalidAccessToken() *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: MessageInvalidAccessToken,
	}
}

func NewAuthenticationErrorDueToRefreshToken() *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: MessageRefreshToken,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}
