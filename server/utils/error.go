package utils

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrJWTNotFound  = errors.New("jwt not found")
	ErrSkip         = errors.New("skip")
	ErrComposeRes   = errors.New("compose response")
	ErrPassword     = errors.New("password error")
	ErrPermDenied   = errors.New("permission denied")
)
