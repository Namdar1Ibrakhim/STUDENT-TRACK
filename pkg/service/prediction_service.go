package service

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
)

type PredictionService struct {
	repo repository.Predictions
}

// Создаем новый сервис
func NewPredictionService(repo repository.Predictions) *PredictionService {
	return &PredictionService{repo: repo}
}

// Получение всех курсов с пагинацией и опциональной сортировкой по оценкам
// func (s *PredictionService) GetAllPredictionByFilter(page, pageSize int, sortBy string) ([]dto.PredictionResponse, error) {
// 	var sortByGrades *string
// 	if sortBy == "asc" || sortBy == "desc" {
// 		sortByGrades = &sortBy
// 	}

// 	return s.repo.GetAllPredictionByFilter(page, pageSize, sortByGrades)
// }
func (s *PredictionService) GetAllPrediction() ([]dto.PredictionResponse, error) {
	return s.repo.GetAllPrediction()
}

// Получение курса студента по его идентификатору
func (s *PredictionService) GetPredictionById(Id int) (dto.PredictionResponse, error) {
	return s.repo.GetPredictionById(Id)
}

// Получение курса студента по идентификатору студента
func (s *PredictionService) GetPredictionByStudentId(studentId int) (dto.PredictionResponse, error) {
	return s.repo.GetPredictionByStudentId(studentId)
}

// Получение курса студента по идентификатору направления
func (s *PredictionService) GetPredictionByDirectionId(directionId int) (dto.PredictionResponse, error) {
	return s.repo.GetPredictionByDirectionId(directionId)
}




