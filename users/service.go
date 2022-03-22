package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input *RegisterInput) (*User, error)
	Login(input *LoginInput) (*User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Register(input *RegisterInput) (*User, error) {
	// cek apakah email sudah dipake
	oldUser, _ := s.repository.FindByEmail(input.Email)
	if oldUser != nil {
		return nil, errors.New("email has been registered")
	}

	encrypt, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	user := User{
		Name:       input.Name,
		Occupation: input.Occupation,
		Email:      input.Email,
		Password:   string(encrypt),
	}
	newUser, err := s.repository.Save(&user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *service) Login(input *LoginInput) (*User, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("email or password wrong")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("email or password wrong")
	}

	return user, nil

}
