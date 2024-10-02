package service

import (
	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	"io"
)

// Все сервисные интерфейсы пишем здесь
type Authorization interface {
	CreateUser(user track.User, role constants.Role) (int, error)
	GenerateToken(username, password string) (string, error)
	//то что принимаем      //то что возвращаем
	ParseToken(token string) (int, error)
	GetUser(userId int) (dto.UserResponse, error)
	UpdateUser(userId int, input dto.UpdateUser) error
	DeleteUser(userId int) error
	EditPassword(userId int, password string) error
}

type CSV interface {
	ValidateCSV(file io.Reader) error
	ProcessCSV(studentId int, file io.Reader) error
}

type Service struct {
	Authorization
	CSV
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CSV:           NewCSVService(repos.CSV),
	}
}
