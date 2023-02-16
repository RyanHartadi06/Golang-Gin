package user

import "gorm.io/gorm"

type RepositoryUser interface {
	FindAll() ([]User, error)
	FindByID(id int) (User, error)
	Create(user User) (User, error)
}

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db}
}

func (r *repositoryUser) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repositoryUser) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Find(&user, ID).Error

	return user, err
}

func (r *repositoryUser) Create(user User) (User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
