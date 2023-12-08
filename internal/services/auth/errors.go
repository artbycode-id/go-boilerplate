package auth

import "artbycode.id/go-app/pkg"

const (
	ErrUserEmailNotFound       = "user email not found"
	CodeErrorUserEmailNotFound = 1001

	ErrInvalidPassword       = "invalid password"
	CodeErrorInvalidPassword = 1002

	TypeAuthError = "auth_error"
)

type AuthErrorDTO struct {
	UniqueId string `json:"unique_id"`
	Type     string `json:"type"`
	Code     int    `json:"code"`
	Message  string `json:"error"`
	Detail   string `json:"detail,omitempty"`
}

func NewAuthError(message string, code int) *AuthErrorDTO {
	return NewAuthErrorWithDetail(message, code, "")
}

func NewAuthErrorWithDetail(message string, code int, detail string) *AuthErrorDTO {
	return &AuthErrorDTO{
		Code:     code,
		Message:  message,
		UniqueId: pkg.GenerateUUID(),
		Type:     TypeAuthError,
		Detail:   detail,
	}
}

func (e *AuthErrorDTO) Error() string {
	return e.Message
}
