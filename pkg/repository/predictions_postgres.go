package repository

import (
	"fmt"
	"time"

	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/jmoiron/sqlx"
)

type PredictionsPostgres struct {
	db *sqlx.DB
}

func NewPredictionsPostgres(db *sqlx.DB) *PredictionsPostgres {
	return &PredictionsPostgres{db: db}
}


func (r *PredictionsPostgres) GetAllPrediction() ([]dto.PredictionResponse, error) {
	query := "SELECT id, student_id, direction_id, created_at FROM prediction"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var predictions []dto.PredictionResponse
	for rows.Next() {
		var prediction dto.PredictionResponse
		err := rows.Scan(&prediction.Id, &prediction.StudentId, &prediction.DirectionId, &prediction.CreatedAt)
		if err != nil {
			return nil, err
		}
		predictions = append(predictions, prediction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return predictions, nil
}

func (r *PredictionsPostgres) SavePrediction(studentId int, directionId int) error {
	_, err := r.db.Exec("INSERT INTO prediction (student_id, direction_id) VALUES ($1, $2)", studentId, directionId)
	if err != nil {
		return fmt.Errorf("failed to save prediction: %v", err)
	}
	return nil
}

func (r *PredictionsPostgres) GetPredictionByStudentId(studentId int) (dto.PredictionResponse, error) {
	var prediction dto.PredictionResponse
	query := "SELECT id, student_id, direction_id, created_at FROM prediction WHERE student_id = $1"

	err := r.db.QueryRow(query, studentId).Scan(&prediction.Id, &prediction.StudentId, &prediction.DirectionId, &prediction.CreatedAt)
	if err != nil {
		return prediction, err
	}

	return prediction, nil
}
func (r *PredictionsPostgres) GetPredictionByDirectionId(directionId int) (dto.PredictionResponse, error) {
	var prediction dto.PredictionResponse
	query := "SELECT id, student_id, direction_id, created_at FROM prediction WHERE direction_id = $1"

	err := r.db.QueryRow(query, directionId).Scan(&prediction.Id, &prediction.StudentId, &prediction.DirectionId, &prediction.CreatedAt)
	if err != nil {
		return prediction, err
	}

	return prediction, nil
}
func (r *PredictionsPostgres) GetPredictionById(id int) (dto.PredictionResponse, error) {
	var prediction dto.PredictionResponse
	query := "SELECT id, student_id, direction_id, created_at FROM prediction WHERE id = $1"

	err := r.db.QueryRow(query, id).Scan(&prediction.Id, &prediction.StudentId, &prediction.DirectionId, &prediction.CreatedAt)
	if err != nil {
		return prediction, err
	}

	return prediction, nil
}

// Function to filter by a date range
func (r *PredictionsPostgres) GetPredictionByDateRange(studentId int, startDate, endDate string) ([]dto.PredictionResponse, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %v", err)
	}

	var predictions []dto.PredictionResponse

	// SQL query to filter by date range
	query := `SELECT * FROM predictions WHERE student_id = $1 AND DATE(created_at) BETWEEN $2 AND $3`
	err = r.db.Select(&predictions, query, studentId, start.Format("2006-01-02"), end.Format("2006-01-02"))

	if err != nil {
		return nil, err
	}

	return predictions, nil
}
