package dto

type PredictionResponseDto struct {
	PredictedTrack string `json:"predicted_track"`
	StudentId      int    `json:"student_id"`
}
