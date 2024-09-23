package dto

type UserResponse struct { // структура пользователя для регистрации
	Id            int    `json:"-" db:"id"`                    // ID не передается в JSON
	Firstname     string `json:"firstname" binding:"required"` // Имя
	Lastname      string `json:"lastname" binding:"required"`  // Имя
	Username      string `json:"username" binding:"required"`  // Имя пользователя
	Password_hash string
	Role          int
}
