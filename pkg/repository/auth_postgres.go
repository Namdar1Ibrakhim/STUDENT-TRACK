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

	checkQuery := fmt.Sprintf("SELECT id FROM %s WHERE email=$1", usersTable)
	err := r.db.QueryRow(checkQuery, user.Email).Scan(&id)
	if err == nil {
		return 0, fmt.Errorf("user with email %s already exists", user.Email)
	}

	insertQuery := fmt.Sprintf("INSERT INTO %s (firstname, lastname, email, password_hash, role) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)

	row := r.db.QueryRow(insertQuery, user.Firstname, user.Lastname, user.Email, user.Password, role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (track.User, error) {
	var user track.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *AuthPostgres) GetAllUsers() ([]dto.GetAllUsersResponse, error) {
	var users []dto.GetAllUsersResponse
	query := fmt.Sprintf("SELECT id, firstname, lastname, email, role FROM %s", usersTable)
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AuthPostgres) GetStudents() ([]dto.StudentsResponse, error) {
	var students []dto.StudentsResponse
	query := fmt.Sprintf("SELECT id, firstname, lastname FROM %s WHERE role = 1", usersTable)
	err := r.db.Select(&students, query)
	return students, err
}

func (r *AuthPostgres) FindByID(userId int) (dto.UserResponse, error) {
	var user dto.UserResponse
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, userId)
	log.Default()
	return user, err
}

func (r *AuthPostgres) UpdateUser(userId int, input dto.UpdateUser) error {
	query := fmt.Sprintf("UPDATE %s SET firstname=$1, lastname=$2, email=$3 WHERE id=$4", usersTable)
	_, err := r.db.Exec(query, input.Firstname, input.Lastname, input.Email, userId)
	return err
}

func (r *AuthPostgres) DeleteUser(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, userId)
	return err
}

func (r *AuthPostgres) GetPasswordHashById(userId int) (dto.GetPasswordRequest, error) {
	var user dto.GetPasswordRequest
	query := fmt.Sprintf("SELECT id, password_hash FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, userId)
	return user, err
}

func (r *AuthPostgres) EditPassword(userId int, password string) error {
	query := fmt.Sprintf("UPDATE %s SET password_hash=$1 WHERE id=$2", usersTable)
	_, err := r.db.Exec(query, password, userId)
	return err
}
