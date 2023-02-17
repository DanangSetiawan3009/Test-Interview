package services

import (
	"kredit-plus/mocks"
	"kredit-plus/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLimitByIdSuccess(t *testing.T) {
	repository := mocks.LimitRepository{}
	repository.On("FindByID", 1).Return(models.LimitKonsumen{}, nil).Once()

	resp, err := NewLimitService(&repository).FindByID(1)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	repository.AssertExpectations(t)
}
