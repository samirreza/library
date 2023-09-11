package service

import (
	"github.com/samirreza/library/internal/model"
	"github.com/samirreza/library/internal/repository"
	"github.com/samirreza/library/internal/service/DTO"
	"github.com/samirreza/library/internal/valueobject"
)

type CreateBookService struct {
	bookRepository repository.BookRepository
}

func NewCreateBookService(br repository.BookRepository) CreateBookService {
	return CreateBookService{bookRepository: br}
}

func (c CreateBookService) Create(cbd DTO.CreateBookDTO) error {
	ulid, err := valueobject.NewULIDFromString(cbd.ID)
	if err != nil {
		return err
	}
	book := model.Book{
		ID:          ulid,
		Title:       cbd.Title,
		Description: cbd.Description,
	}
	return c.bookRepository.Add(book)
}
