package book

import (
	"gorm.io/gorm"
)

type RepositoryBook interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book Book) (Book, error)
}

type repositoryBook struct {
	db *gorm.DB
}

func NewRepositoryBook(db *gorm.DB) *repositoryBook {
	return &repositoryBook{db}
}

func (r *repositoryBook) FindAll() ([]Book, error){
	var books []Book

	err := r.db.Find(&books).Error


	return books, err
}

func (r *repositoryBook) FindByID(ID int) (Book, error){
	var book Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repositoryBook) Create(book Book) (Book, error){
	err := r.db.Create(&book).Error

	return book, err
}