package repositories

import (
	"kredit-plus/models"

	"gorm.io/gorm"
)

type LimitRepository interface {
	FindAll() ([]models.LimitKonsumen, error)
	FindByID(ID int) (models.LimitKonsumen, error)
	Create(limit models.LimitKonsumen) (models.LimitKonsumen, error)
	Update(limit models.LimitKonsumen) (models.LimitKonsumen, error)
	Delete(limit models.LimitKonsumen) (models.LimitKonsumen, error)
}

type limitRepository struct {
	db *gorm.DB
}

func NewLimitRepository(db *gorm.DB) *limitRepository {
	return &limitRepository{db}
}

func (r *limitRepository) FindAll() ([]models.LimitKonsumen, error) {
	limit := []models.LimitKonsumen{}
	result := r.db.Find(&limit)

	return limit, result.Error
}

func (r *limitRepository) FindByID(ID int) (models.LimitKonsumen, error) {
	var limit models.LimitKonsumen
	err := r.db.Find(&limit, ID).Error

	return limit, err
}

func (r *limitRepository) Create(model models.LimitKonsumen) (models.LimitKonsumen, error) {
	err := r.db.Create(&model).Error

	return model, err
}

func (r *limitRepository) Update(model models.LimitKonsumen) (models.LimitKonsumen, error) {
	err := r.db.Save(&model).Error

	return model, err
}

func (r *limitRepository) Delete(model models.LimitKonsumen) (models.LimitKonsumen, error) {
	err := r.db.Delete(&model).Error

	return model, err
}
