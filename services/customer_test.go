package services

import (
	"kredit-plus/mocks"
	"kredit-plus/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCustomerByIdSuccess(t *testing.T) {
	repository := mocks.CustomerRepository{}
	repository.On("FindByID", 1).Return(models.Konsumen{}, nil).Once()

	resp, err := NewCustomerService(&repository).FindByID(1)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	repository.AssertExpectations(t)
}
