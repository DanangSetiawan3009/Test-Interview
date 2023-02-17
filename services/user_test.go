package services

import (
	"errors"
	"kredit-plus/mocks"
	"kredit-plus/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisUserSuccess(t *testing.T) {
	model := models.User{}
	password, _ := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	model.Password = string(password)

	repository := mocks.UserRepository{}
	repository.On("FindByEmail", model.Email).Return(nil, errors.New("")).Once()
	repository.On("Regis", model).Return(models.User{}, nil).Once()

	err := NewUserService(&repository).Regis(model)

	assert.Nil(t, err)
	repository.AssertExpectations(t)
}

func TestRegisError(t *testing.T) {
	model := models.User{}

	repository := mocks.UserRepository{}
	repository.On("FindByEmail", model.Email).Return(errors.New("email already in use")).Once()

	resp, err := NewUserService(&repository).userRepo.FindByEmail(model.Email)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	repository.AssertExpectations(t)
}
