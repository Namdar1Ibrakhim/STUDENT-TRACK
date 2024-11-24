package dto

type StudentsResponse struct {
	Id            int    `json:"id"`
	Firstname     string `json:"firstname" binding:"required"`
	Lastname      string `json:"lastname" binding:"required"`
	Email         string `json:"-" db:"email"`
	Password_hash string `json:"-" db:"password_hash"`
	Role          int    `json:"-" db:"Role"`
}
