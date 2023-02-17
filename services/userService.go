package services

import (
	"fmt"
	"kredit-plus/models"
	"kredit-plus/repositories"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Regis(userRegis models.User) error
	Login(userLogin models.User) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepo}
}

func (s *userService) Regis(userRegis models.User) error {
	_, err := s.userRepo.FindByEmail(userRegis.Email)
	if err == nil {
		return fmt.Errorf("email already in use")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(userRegis.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	user := &models.User{
		Fullname: userRegis.Fullname,
		Email:    userRegis.Email,
		Password: string(password),
	}

	err = s.userRepo.Regis(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Login(userLogin models.User) (string, error) {
	user, err := s.userRepo.FindByEmail(userLogin.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token, err := createToken(userLogin.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("secret"))
}
