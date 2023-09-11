package repository

import (
	"github.com/samirreza/library/internal/model"
	"github.com/samirreza/library/internal/valueobject"
)

type BookRepositoryInterface interface {
	Get(valueobject.ULID) (model.Book, error)
	Add(model.Book) error
	Update(model.Book) error
	Delete(model.Book) error
}
