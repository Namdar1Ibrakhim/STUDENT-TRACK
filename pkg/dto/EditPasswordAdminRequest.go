package dto

type EditPasswordAdminRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}
