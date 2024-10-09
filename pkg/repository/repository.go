package repository

import (
	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/jmoiron/sqlx"
)

// Все репозиторные интерфейсы пишем здесь
type Authorization interface {
	CreateUser(user track.User, role constants.Role) (int, error)
	GetUser(username, password string) (track.User, error)
	FindByID(userId int) (dto.UserResponse, error)
	UpdateUser(userId int, input dto.UpdateUser) error
	DeleteUser(userId int) error
	EditPassword(userId int, password string) error
}

type Predictions interface {
	SavePrediction(userId int, predictions string) error
}
type Course interface {
	GetAllCourse() ([]dto.CourseResponse, error)
	GetCourseById(courseId int) (dto.CourseResponse, error)
	GetCourseByName(courseName string) (dto.CourseResponse, error)
}

type Direction interface {
	GetAllDirection() ([]dto.DirectionResponse, error)
	GetDirectionById(directionId int) (dto.DirectionResponse, error)
	GetDirectionByName(directionName string) (dto.DirectionResponse, error)
}

type StudentCourse interface {
	GetAllStudentCourse() ([]dto.StudentCourseResponse, error)
	GetStudentCourseById(studentCourseId int) (dto.StudentCourseResponse, error)
	GetStudentCourseByStudentId(studentId int) (dto.StudentCourseResponse, error)
	GetStudentCourseByCourseId(courseId int) (dto.StudentCourseResponse, error)
	GetAllStudentCourseByFilter(limit, page int, sortByGrades *string) ([]dto.StudentCourseResponse, error)
	AddStudentCourse(student_id int, course_id int, grades int) error
}

type Repository struct {
	Authorization
	Predictions
	Course
	Direction
	StudentCourse
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Predictions:   NewPredictionsPostgres(db),
		Course:        NewCourseRepository(db),
		Direction:     NewDirectionRepository(db),
		StudentCourse: NewStudentCoursePostgres(db),
	}
}
