package dto

type UpdateUser struct {
	Id        int    `json:"-" db:"id"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required"`
}
