package constants

import "errors"

// Ошибки, связанные с пользователями
var (
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidUserId        = errors.New("invalid user id")
	ErrInvalidStudentId     = errors.New("invalid student id")
	ErrStudentNotFound      = errors.New("student not found")
	ErrIncorrectOldPassword = errors.New("incorrect old password")
)

// Ошибки, связанные с файлами
var (
	ErrFileNotFound        = errors.New("file not found")
	ErrFailedToUploadFile  = errors.New("failed to upload file")
	ErrFailedToOpenFile    = errors.New("failed to open file")
	ErrFailedToProcessFile = errors.New("failed to process file")
	ErrFailedToReadFile    = errors.New("failed to read file")
	ErrInvalidCSVStructure = errors.New("invalid CSV structure")
)

// Ошибки доступа и авторизации
var (
	ErrAccessDenied         = errors.New("you don't have access to this resource")
	ErrEmptyAuthHeader      = errors.New("empty authorization header")
	ErrInvalidAuthHeader    = errors.New("invalid authorization header")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
)

// Ошибки, связанные с сервисами
var (
	ErrMLServiceFailure = errors.New("ML service prediction failed")
)

// Общие ошибки ввода данных
var (
	ErrInvalidInputData = errors.New("invalid input data")
)
