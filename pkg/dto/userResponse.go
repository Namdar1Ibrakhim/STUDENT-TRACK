package dto

type UserResponse struct {
	Id            int    `json:"-" db:"id"`                   
	Firstname     string `json:"firstname" binding:"required"` 
	Lastname      string `json:"lastname" binding:"required"`  
	Username      string `json:"username" binding:"required"`  
	Password_hash string
	Role          int
}
