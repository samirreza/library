package service

import (
	"github.com/samirreza/library/internal/repository"
	"github.com/samirreza/library/internal/service/DTO"
	"github.com/samirreza/library/internal/valueobject"
)

type UpdateBookService struct {
	bookRepository repository.BookRepository
}

func NewUpdateBookService(br repository.BookRepository) UpdateBookService {
	return UpdateBookService{bookRepository: br}
}

func (c UpdateBookService) Update(id string, ubd DTO.UpdateBookDTO) error {
	ulid, err := valueobject.NewULIDFromString(id)
	if err != nil {
		return err
	}
	book, err := c.bookRepository.Get(ulid)
	if err != nil {
		return err
	}
	book.Title = ubd.Title
	book.Description = ubd.Description
	return c.bookRepository.Update(book)
}
