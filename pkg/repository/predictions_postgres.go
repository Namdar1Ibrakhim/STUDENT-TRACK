package repository

import (
	"github.com/jmoiron/sqlx"
)

type PredictionsPostgres struct {
	db *sqlx.DB
}

func NewPredictionsPostgres(db *sqlx.DB) *PredictionsPostgres {
	return &PredictionsPostgres{db: db}
}

func (r *PredictionsPostgres) SavePrediction(userId int, predictions string) error {
	_, err := r.db.Exec("INSERT INTO prediction (student_id, prediction_text) VALUES ($1, $2)", userId, predictions)
	return err
}
