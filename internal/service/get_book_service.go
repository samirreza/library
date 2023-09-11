package service

import (
	"github.com/samirreza/library/internal/repository"
	"github.com/samirreza/library/internal/service/DTO"
	"github.com/samirreza/library/internal/valueobject"
)

type GetBookService struct {
	bookRepository repository.BookRepository
}

func NewGetBookService(br repository.BookRepository) GetBookService {
	return GetBookService{bookRepository: br}
}

func (g GetBookService) Get(id string) (DTO.GetBookDTO, error) {
	ulid, err := valueobject.NewULIDFromString(id)
	if err != nil {
		return DTO.GetBookDTO{}, err
	}
	book, err := g.bookRepository.Get(ulid)
	return DTO.GetBookDTO{
		ID:          book.ID.String(),
		Title:       book.Title,
		Description: book.Description,
	}, err
}
