package repositories

import (
	"kredit-plus/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	FindAll() ([]models.Transaksi, error)
	FindByID(ID int) (models.Transaksi, error)
	Create(model models.Transaksi) (models.Transaksi, error)
	Update(model models.Transaksi) (models.Transaksi, error)
	Delete(model models.Transaksi) (models.Transaksi, error)
}

type transaksiRepository struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *transaksiRepository {
	return &transaksiRepository{db}
}

func (r *transaksiRepository) FindAll() ([]models.Transaksi, error) {
	transaksi := []models.Transaksi{}
	result := r.db.Find(&transaksi)

	return transaksi, result.Error
}

func (r *transaksiRepository) FindByID(ID int) (models.Transaksi, error) {
	var transaksi models.Transaksi
	err := r.db.Find(&transaksi, ID).Error

	return transaksi, err
}

func (r *transaksiRepository) Create(model models.Transaksi) (models.Transaksi, error) {
	err := r.db.Create(&model).Error

	return model, err
}

func (r *transaksiRepository) Update(model models.Transaksi) (models.Transaksi, error) {
	err := r.db.Save(&model).Error

	return model, err
}

func (r *transaksiRepository) Delete(model models.Transaksi) (models.Transaksi, error) {
	err := r.db.Delete(&model).Error

	return model, err
}
