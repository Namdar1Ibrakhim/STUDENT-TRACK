package repository

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/jmoiron/sqlx"
)

type DirectionRepository struct {
	db *sqlx.DB
}

func NewDirectionRepository(db *sqlx.DB) *DirectionRepository {
	return &DirectionRepository{db: db}
}

func (r *DirectionRepository) GetAllDirection() ([]dto.DirectionResponse, error) {
	query := "SELECT id, direction_name, description FROM direction"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var directions []dto.DirectionResponse
	for rows.Next() {
		var direction dto.DirectionResponse
		err := rows.Scan(&direction.Id, &direction.Direction_name, &direction.Description)
		if err != nil {
			return nil, err
		}
		directions = append(directions, direction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return directions, nil
}

func (r *DirectionRepository) GetDirectionById(directionId int) (dto.DirectionResponse, error) {
	var direction dto.DirectionResponse
	query := "SELECT id, direction_name, description FROM direction WHERE id = $1"

	err := r.db.QueryRow(query, directionId).Scan(&direction.Id, &direction.Direction_name, &direction.Description)
	if err != nil {
		return direction, err
	}

	return direction, nil
}

func (r *DirectionRepository) GetDirectionByName(directionName string) (dto.DirectionResponse, error) {
	var direction dto.DirectionResponse
	query := "SELECT id, direction_name, description FROM direction WHERE direction_name = $1"

	err := r.db.QueryRow(query, directionName).Scan(&direction.Id, &direction.Direction_name, &direction.Description)
	if err != nil {
		return direction, err
	}

	return direction, nil
}
