package services

import (
	"kredit-plus/models"
	"kredit-plus/repositories"
)

type CustomerService interface {
	FindByID(ID int) (models.Konsumen, error)
	GetAll() ([]models.Konsumen, error)
	Create(cust models.Konsumen) (models.Konsumen, error)
	Update(ID int, cust models.Konsumen) (models.Konsumen, error)
	Delete(ID int) (models.Konsumen, error)
}

type service struct {
	repository repositories.CustomerRepository
}

func NewCustomerService(repository repositories.CustomerRepository) *service {
	return &service{repository}
}

// service for konsumen
func (s *service) Create(model models.Konsumen) (models.Konsumen, error) {
	cust := models.Konsumen{
		Nik:          model.Nik,
		Fullname:     model.Fullname,
		Legalname:    model.Legalname,
		TempatLahir:  model.TempatLahir,
		TanggalLahir: model.TanggalLahir,
		Gaji:         model.Gaji,
		FotoKtp:      model.FotoKtp,
		FotoSelfie:   model.FotoSelfie,
	}
	newCust, err := s.repository.Create(cust)

	return newCust, err
}

func (s *service) GetAll() ([]models.Konsumen, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (models.Konsumen, error) {
	cust, err := s.repository.FindByID(ID)

	return cust, err
}

func (s *service) Update(ID int, request models.Konsumen) (models.Konsumen, error) {
	cust, err := s.repository.FindByID(ID)
	if err != nil {
		return cust, err
	}

	cust.Nik = request.Nik
	cust.Fullname = request.Fullname

	newCust, err := s.repository.Update(cust)
	if err != nil {
		return newCust, err
	}
	return newCust, err
}

func (s *service) Delete(ID int) (models.Konsumen, error) {
	cust, err := s.repository.FindByID(ID)
	if err != nil {
		return cust, err
	}

	newCust, err := s.repository.Delete(cust)
	return newCust, err
}
