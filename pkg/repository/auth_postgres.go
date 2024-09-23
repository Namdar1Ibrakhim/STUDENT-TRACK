package repository

import (
	"fmt"
	"log"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user track.User, role constants.Role) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, username, password_hash, role) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Firstname, user.Lastname, user.Username, user.Password, role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (track.User, error) {
	var user track.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) FindByID(userId int) (dto.UserResponse, error) {
	var user dto.UserResponse
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, userId)
	log.Default()
	return user, err
}
