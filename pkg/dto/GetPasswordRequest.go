package dto

type GetPasswordRequest struct {
	Id            int    `json:"id"`
	Password_hash string `json:"password_hash" db:"password_hash"`
}
