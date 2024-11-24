package track

type User struct { // структура пользователя для регистрации
	Id        int    `json:"-" db:"id"`                    // ID не передается в JSON
	Firstname string `json:"firstname" binding:"required"` // Имя
	Lastname  string `json:"lastname" binding:"required"`  // Имя
	Email     string `json:"email" binding:"required"`  // Имя пользователя
	Password  string `json:"password" binding:"required"`  // Пароль
}
