package dto

type GetAllUsersResponse struct {
	Id            int    `json:"id"`
	Firstname     string `json:"firstname" binding:"required"`
	Lastname      string `json:"lastname" binding:"required"`
	Username      string `json:"username" binding:"required"`
	Password_hash string `json:"-" db:"password_hash"`
	Role          int
}
