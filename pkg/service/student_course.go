package service

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
)

type StudentCourseService struct {
	repo repository.StudentCourse
}

// Создаем новый сервис
func NewStudentCourseService(repo repository.StudentCourse) *StudentCourseService {
	return &StudentCourseService{repo: repo}
}

// Получение всех курсов с пагинацией и опциональной сортировкой по оценкам
func (s *StudentCourseService) GetAllStudentCourseByFilter(page, pageSize int, sortBy string) ([]dto.StudentCourseResponse, error) {
	var sortByGrades *string
	if sortBy == "asc" || sortBy == "desc" {
		sortByGrades = &sortBy
	}

	return s.repo.GetAllStudentCourseByFilter(page, pageSize, sortByGrades)
}
func (s *StudentCourseService) GetAllStudentCourse() ([]dto.StudentCourseResponse, error) {
	return s.repo.GetAllStudentCourse()
}

// Получение курса студента по его идентификатору
func (s *StudentCourseService) GetStudentCourseById(courseId int) (dto.StudentCourseResponse, error) {
	return s.repo.GetStudentCourseById(courseId)
}

// Получение курса студента по идентификатору студента
func (s *StudentCourseService) GetStudentCourseByStudentId(studentId int) (dto.StudentCourseResponse, error) {
	return s.repo.GetStudentCourseByStudentId(studentId)
}

// Получение курса студента по идентификатору курса
func (s *StudentCourseService) GetStudentCourseByCourseId(courseId int) (dto.StudentCourseResponse, error) {
	return s.repo.GetStudentCourseByCourseId(courseId)
}
