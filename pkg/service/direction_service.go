package service

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
)

type DirectionService struct {
	repo repository.Direction
}

func NewDirectionService(repo repository.Direction) *DirectionService {
	return &DirectionService{repo: repo}
}

func (s *DirectionService) GetAllDirection() ([]dto.DirectionResponse, error) {
    return s.repo.GetAllDirection()
}

func (s *DirectionService) GetDirectionById(directionId int) (dto.DirectionResponse, error) {
    return s.repo.GetDirectionById(directionId)
}

func (s *DirectionService) GetDirectionByName(directionName string) (dto.DirectionResponse, error) {
    return s.repo.GetDirectionByName(directionName)
}