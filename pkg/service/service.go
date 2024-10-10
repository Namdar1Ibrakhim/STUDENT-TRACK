package service

import (
	"io"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
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
	PredictCSV(studentId int, file io.Reader) (string, error)
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
	GetAllStudentCourseByFilter(pageSize int, page int, sortByGrades string) ([]dto.StudentCourseResponse, error)
	AddStudentCourse(student_id int, course_id int, grades int) error
}

type Service struct {
	Authorization
	CSV
	Course
	Direction
	StudentCourse
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CSV:           NewCSVService(repos.Predictions, repos.Course, repos.StudentCourse, repos.Direction),
		Course:        NewCourseService(repos.Course),
		Direction:     NewDirectionService(repos.Direction),
		StudentCourse: NewStudentCourseService(repos.StudentCourse),
	}
}
