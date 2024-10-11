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

<<<<<<< HEAD
func (r *PredictionsPostgres) SavePrediction(studentId int, directionId int) error {
	_, err := r.db.Exec("INSERT INTO prediction (student_id, direction_id) VALUES ($1, $2)", studentId, directionId)
	return err
=======
func (r *PredictionsPostgres) SavePrediction(userId int, directionId int) error {
	_, err := r.db.Exec("INSERT INTO prediction (student_id, direction_id) VALUES ($1, $2)", userId, directionId)
	if err != nil {
		return fmt.Errorf("failed to save prediction: %v", err)
	}
	return nil
>>>>>>> e1e18e5e99ee210f33fd65ee2b1bb3d695728391
}
