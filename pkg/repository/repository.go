package repository

import (
	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/jmoiron/sqlx"
)

//Все репозиторные интерфейсы пишем здесь
type Authorization interface {
	CreateUser(user track.User) (int, error)
	GetUser(username, password string) (track.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
