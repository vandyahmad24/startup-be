package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input *RegisterInput) (*User, error)
	Login(input *LoginInput) (*User, error)
	CheckEmail(input *CheckEmailInput) (*User, error)
	SaveAvatar(ID int, fileLocation string) (*User, error)
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

func (s *service) CheckEmail(input *CheckEmailInput) (*User, error) {

	user, _ := s.repository.FindByEmail(input.Email)
	if user != nil {
		return nil, errors.New("email has been registered")
	}
	return user, nil

}

func (s *service) SaveAvatar(ID int, fileLocation string) (*User, error) {
	oldUser, err := s.repository.FindById(ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	oldUser.Avatar = fileLocation
	newUser, err := s.repository.Update(oldUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}
