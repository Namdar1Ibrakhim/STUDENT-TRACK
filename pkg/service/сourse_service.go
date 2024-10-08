package service

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetAllCourse() ([]dto.CourseResponse, error) {
    return s.repo.GetAll()
}

func (s *CourseService) GetCourseById(courseId int) (dto.CourseResponse, error) {
    return s.repo.GetById(courseId)
}

func (s *CourseService) GetCourseByName(courseName string) (dto.CourseResponse, error) {
    return s.repo.GetByName(courseName)
}