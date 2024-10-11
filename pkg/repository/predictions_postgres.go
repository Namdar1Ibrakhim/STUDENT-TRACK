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

func (r *PredictionsPostgres) SavePrediction(studentId int, directionId int) error {
	_, err := r.db.Exec("INSERT INTO prediction (student_id, direction_id) VALUES ($1, $2)", studentId, directionId)
	return err
}
