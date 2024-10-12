package dto

type PredictionResponse struct {
	Id            int     `json:"-" db:"id"`                   
	StudentId     int     `json:"student_id"`
	DirectionId   int     `json:"direction_id"`
	CreatedAt     string  `json:"created_at" db:"created_at"`
}
