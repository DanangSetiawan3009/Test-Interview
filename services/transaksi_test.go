package services

import (
	"kredit-plus/mocks"
	"kredit-plus/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTransaksiByIdSuccess(t *testing.T) {
	repository := mocks.TransaksiRepository{}
	repository.On("FindByID", 1).Return(models.Transaksi{}, nil).Once()

	resp, err := NewTransaksiService(&repository).FindByID(1)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	repository.AssertExpectations(t)
}
