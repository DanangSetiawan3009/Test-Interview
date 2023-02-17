package repositories

import (
	"kredit-plus/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll() ([]models.Konsumen, error)
	FindByID(ID int) (models.Konsumen, error)
	Create(cust models.Konsumen) (models.Konsumen, error)
	Update(cust models.Konsumen) (models.Konsumen, error)
	Delete(cust models.Konsumen) (models.Konsumen, error)
}

type repository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Konsumen, error) {
	customers := []models.Konsumen{}
	result := r.db.Find(&customers)

	return customers, result.Error
}

func (r *repository) FindByID(ID int) (models.Konsumen, error) {
	var cust models.Konsumen
	err := r.db.Find(&cust, ID).Error

	return cust, err
}

func (r *repository) Create(cust models.Konsumen) (models.Konsumen, error) {
	err := r.db.Create(&cust).Error

	return cust, err
}

func (r *repository) Update(cust models.Konsumen) (models.Konsumen, error) {
	err := r.db.Save(&cust).Error

	return cust, err
}

func (r *repository) Delete(cust models.Konsumen) (models.Konsumen, error) {
	err := r.db.Delete(&cust).Error

	return cust, err
}
