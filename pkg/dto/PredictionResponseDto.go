package dto

type PredictionResponseDto struct {
	Id             int    `json:"-" db:"id"`
	Direction_name string `json:"direction_name" binding:"required"`
	StudentId      int    `json:"student_id"`
}
