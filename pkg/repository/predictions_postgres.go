package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PredictionsPostgres struct {
	db *sqlx.DB
}

func NewPredictionsPostgres(db *sqlx.DB) *PredictionsPostgres {
	return &PredictionsPostgres{db: db}
}

func (r *PredictionsPostgres) SavePrediction(userId int, directionId int) error {
	_, err := r.db.Exec("INSERT INTO prediction (student_id, direction_id) VALUES ($1, $2)", userId, directionId)
	if err != nil {
		return fmt.Errorf("failed to save prediction: %v", err)
	}
	return nil
}
