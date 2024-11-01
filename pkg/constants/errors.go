package constants

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrFileNotFound      = errors.New("file not found")
	ErrAccessDenied      = errors.New("you don't have access to this resource")
	ErrEmptyAuthHeader   = errors.New("empty authorization header")
	ErrInvalidAuthHeader = errors.New("invalid authorization header")
)
var (
	ErrInvalidCSVStructure = errors.New("invalid CSV structure")
	ErrMLServiceFailure    = errors.New("ML service prediction failed")
)
