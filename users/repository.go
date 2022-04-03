package users

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user *User) (*User, error)
	FindByEmail(email string) (*User, error)
	FindById(id int) (*User, error)
	Update(user *User) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(user *User) (*User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindById(id int) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Update(user *User) (*User, error) {

	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}
