package users

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input *RegisterInput) (*User, error)
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
