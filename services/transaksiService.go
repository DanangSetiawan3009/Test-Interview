package services

import (
	"kredit-plus/models"
	"kredit-plus/repositories"
)

type TransaksiService interface {
	FindByID(ID int) (models.Transaksi, error)
	GetAll() ([]models.Transaksi, error)
	Create(model models.Transaksi) (models.Transaksi, error)
	Update(ID int, model models.Transaksi) (models.Transaksi, error)
	Delete(ID int) (models.Transaksi, error)
}

type transaksiService struct {
	transaksiRepo repositories.TransaksiRepository
}

func NewTransaksiService(transaksiRepo repositories.TransaksiRepository) *transaksiService {
	return &transaksiService{transaksiRepo}
}

func (s *transaksiService) Create(model models.Transaksi) (models.Transaksi, error) {
	transaksi := models.Transaksi{
		LimitKonsumenId: model.LimitKonsumenId,
		NoKontrak:       model.NoKontrak,
		Otr:             model.Otr,
		AdminFee:        model.AdminFee,
		JumlahCicilan:   model.JumlahCicilan,
		JumlahBunga:     model.JumlahBunga,
		NamaAsset:       model.NamaAsset,
	}
	newTransaksi, err := s.transaksiRepo.Create(transaksi)

	return newTransaksi, err
}

func (s *transaksiService) GetAll() ([]models.Transaksi, error) {
	return s.transaksiRepo.FindAll()
}

func (s *transaksiService) FindByID(ID int) (models.Transaksi, error) {
	cust, err := s.transaksiRepo.FindByID(ID)

	return cust, err
}

func (s *transaksiService) Update(ID int, request models.Transaksi) (models.Transaksi, error) {
	transaksi, err := s.transaksiRepo.FindByID(ID)
	if err != nil {
		return transaksi, err
	}

	transaksi.NoKontrak = request.NoKontrak
	transaksi.Otr = request.Otr
	transaksi.AdminFee = request.AdminFee
	transaksi.JumlahCicilan = request.JumlahCicilan
	transaksi.JumlahBunga = request.JumlahBunga
	transaksi.NamaAsset = request.NamaAsset

	newTransaksi, err := s.transaksiRepo.Update(transaksi)
	if err != nil {
		return newTransaksi, err
	}
	return newTransaksi, err
}

func (s *transaksiService) Delete(ID int) (models.Transaksi, error) {
	transaksi, err := s.transaksiRepo.FindByID(ID)
	if err != nil {
		return transaksi, err
	}

	newTransaksi, err := s.transaksiRepo.Delete(transaksi)
	return newTransaksi, err
}
