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
	GetUser(email, password string) (track.User, error)
	GetAllUsers() ([]dto.GetAllUsersResponse, error)
	GetStudents() ([]dto.StudentsResponse, error)
	FindByID(userId int) (dto.UserResponse, error)
	UpdateUser(userId int, input dto.UpdateUser) error
	DeleteUser(userId int) error
	GetPasswordHashById(userId int) (dto.GetPasswordRequest, error)
	EditPassword(userId int, newPasswordHash string) error
}

type Predictions interface {
	SavePrediction(studentId int, directionId int) error
	GetAllPrediction() ([]dto.PredictionResponse, error)
	GetPredictionById(id int) (dto.PredictionResponse, error)
	GetPredictionByStudentId(studentId int) (dto.PredictionResponse, error)
	GetPredictionByDirectionId(directionId int) (dto.PredictionResponse, error)
	GetPredictionByDateRange(studentId int, startDate, endDate string) ([]dto.PredictionResponse, error)
}
type Course interface {
	GetAllCourse() ([]dto.CourseResponse, error)
	GetCourseById(courseId int) (dto.CourseResponse, error)
	GetCourseByName(courseName string) (dto.CourseResponse, error)
	FindCourseIDByName(courseName string) (int, error)
}

type Direction interface {
	GetAllDirection() ([]dto.DirectionResponse, error)
	GetDirectionById(directionId int) (dto.DirectionResponse, error)
	GetDirectionByName(directionName string) (dto.DirectionResponse, error)
	FindDirectionIDByName(directionName string) (int, error)
}

type StudentCourse interface {
	GetAllStudentCourse() ([]dto.StudentCourseResponse, error)
	GetStudentCourseById(studentCourseId int) (dto.StudentCourseResponse, error)
	GetStudentCourseByStudentId(studentId int) (dto.StudentCourseResponse, error)
	GetStudentCourseByCourseId(courseId int) (dto.StudentCourseResponse, error)
	GetAllStudentCourseByFilter(limit, page int, sortByGrades *string) ([]dto.StudentCourseResponse, error)
	AddStudentCourse(student_id int, course_id int, grades int) error
}

// type Prediction

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
