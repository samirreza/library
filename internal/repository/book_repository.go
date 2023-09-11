package repository

import (
	"errors"

	"github.com/samirreza/library/internal/model"
	"github.com/samirreza/library/internal/valueobject"
	"gorm.io/gorm"
)

type BookRepository struct {
	db gorm.DB
}

func NewBookRepository(db gorm.DB) BookRepository {
	return BookRepository{db: db}
}

func (b BookRepository) Get(id valueobject.ULID) (model.Book, error) {
	book := model.Book{}
	if result := b.db.First(&book, "id = ?", id.String()); result.Error != nil {
		return book, errors.New("book not found")
	}
	return book, nil
}

func (b BookRepository) Add(book model.Book) error {
	result := b.db.Create(&book)
	return result.Error
}

func (b BookRepository) Update(book model.Book) error {
	result := b.db.Save(&book)
	return result.Error
}

func (b BookRepository) Delete(book model.Book) error {
	result := b.db.Delete(&book)
	return result.Error
}
