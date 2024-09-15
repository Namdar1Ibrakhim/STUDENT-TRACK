package repository

import (
	"fmt"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user track.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, username, password_hash) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Firstname, user.Lastname, user.Username, user.Password)
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
