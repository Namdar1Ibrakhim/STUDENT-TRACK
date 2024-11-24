package dto

type UserResponse struct {
	Id            int    `json:"-" db:"id"`
	Firstname     string `json:"firstname" binding:"required"`
	Lastname      string `json:"lastname" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Password_hash string `json:"-" db:"password_hash"`
	Role          int
}
