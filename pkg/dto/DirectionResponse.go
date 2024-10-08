package dto

type DirectionResponse struct {
    Id            int     `json:"-" db:"id"`                   
	Direction_name   string  `json:"direction_name" binding:"required"`
	Description   string  `json:"description" binding:"required"`

}