package service

import (
	"github.com/samirreza/library/internal/repository"
	"github.com/samirreza/library/internal/valueobject"
)

type DeleteBookService struct {
	bookRepository repository.BookRepository
}

func NewDeleteBookService(br repository.BookRepository) DeleteBookService {
	return DeleteBookService{bookRepository: br}
}

func (g DeleteBookService) Delete(id string) error {
	ulid, err := valueobject.NewULIDFromString(id)
	if err != nil {
		return err
	}
	book, err := g.bookRepository.Get(ulid)
	if err != nil {
		return err
	}
	return g.bookRepository.Delete(book)
}
