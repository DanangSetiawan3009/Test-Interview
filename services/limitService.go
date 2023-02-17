package services

import (
	"kredit-plus/models"
	"kredit-plus/repositories"
)

type LimitService interface {
	GetAll() ([]models.LimitKonsumen, error)
	FindByID(ID int) (models.LimitKonsumen, error)
	Create(limit models.LimitKonsumen) (models.LimitKonsumen, error)
	Update(ID int, limit models.LimitKonsumen) (models.LimitKonsumen, error)
	Delete(ID int) (models.LimitKonsumen, error)
}

type limitService struct {
	limitRepo repositories.LimitRepository
}

func NewLimitService(limitRepo repositories.LimitRepository) *limitService {
	return &limitService{limitRepo}
}

// service for LimitKonsumen
func (s *limitService) Create(model models.LimitKonsumen) (models.LimitKonsumen, error) {
	limit := models.LimitKonsumen{
		KonsumenId: model.KonsumenId,
		Tenor1:     model.Tenor1,
		Tenor2:     model.Tenor2,
		Tenor3:     model.Tenor3,
		Tenor4:     model.Tenor4,
	}
	newLimit, err := s.limitRepo.Create(limit)

	return newLimit, err
}

func (s *limitService) GetAll() ([]models.LimitKonsumen, error) {
	return s.limitRepo.FindAll()
}

func (s *limitService) FindByID(ID int) (models.LimitKonsumen, error) {
	limit, err := s.limitRepo.FindByID(ID)

	return limit, err
}

func (s *limitService) Update(ID int, request models.LimitKonsumen) (models.LimitKonsumen, error) {
	limit, err := s.limitRepo.FindByID(ID)
	if err != nil {
		return limit, err
	}

	limit.Tenor1 = request.Tenor1
	limit.Tenor2 = request.Tenor2
	limit.Tenor3 = request.Tenor3
	limit.Tenor4 = request.Tenor4

	newLimit, err := s.limitRepo.Update(limit)
	if err != nil {
		return newLimit, err
	}
	return newLimit, err
}

func (s *limitService) Delete(ID int) (models.LimitKonsumen, error) {
	limit, err := s.limitRepo.FindByID(ID)
	if err != nil {
		return limit, err
	}

	newLimit, err := s.limitRepo.Delete(limit)
	return newLimit, err
}
