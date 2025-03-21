package service

import (
	"io"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	pb "github.com/Namdar1Ibrakhim/student-track-system/proto"
)

// Все сервисные интерфейсы пишем здесь
type Authorization interface {
	CreateUser(user track.User, role constants.Role) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUser(userId int) (dto.UserResponse, error)
	GetAllUsers() ([]dto.GetAllUsersResponse, error)
	UpdateUser(userId int, input dto.UpdateUser) error
	DeleteUser(userId int) error
	EditPassword(userId int, oldPassword, newPassword string, isAdmin bool) error
	GetStudents() ([]dto.StudentsResponse, error)
}

type CSV interface {
	ValidateCSVForStudent(file io.Reader) error
	ValidateCSVForInstructor(file io.Reader) error
	PredictCSV(studentId int, file io.Reader, isInstructor bool) (map[int]*dto.PredictionResponseDto, error)
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

type Prediction interface {
	GetAllPrediction() ([]dto.PredictionResponse, error)
	GetPredictionById(studentCourseId int) (dto.PredictionResponse, error)
	GetPredictionByStudentId(studentId int) (dto.PredictionResponse, error)
	GetPredictionByDirectionId(directionId int) (dto.PredictionResponse, error)
	//GetAllPredictionByFilter(pageSize int, page int, sortByGrades string) ([]dto.PredictionResponse, error)
}

type Service struct {
	Authorization
	CSV
	Course
	Direction
	StudentCourse
	Prediction
}

func NewService(repos *repository.Repository, mlClient pb.PredictionServiceClient) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CSV:           NewCSVService(repos.Predictions, repos.Course, repos.StudentCourse, repos.Direction, mlClient),
		Course:        NewCourseService(repos.Course),
		Direction:     NewDirectionService(repos.Direction),
		StudentCourse: NewStudentCourseService(repos.StudentCourse),
		Prediction:    NewPredictionService(repos.Predictions),
	}
}
