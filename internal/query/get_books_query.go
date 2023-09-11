package query

import (
	"github.com/samirreza/library/internal/model"
	"gorm.io/gorm"
)

type GetBooksQuery struct {
	db gorm.DB
}

func NewGetBooksQuery(db gorm.DB) GetBooksQuery {
	return GetBooksQuery{db: db}
}

func (g GetBooksQuery) Get(gbi GetBooksInput) ([]GetBookOutput, error) {
	books := []model.Book{}
	query := g.db.Scopes(paginate(gbi.PaginationInput))
	if len(gbi.Title) > 0 {
		query.Where("title LIKE ?", "%"+gbi.Title+"%")
	}
	if len(gbi.Description) > 0 {
		query.Where("description LIKE ?", "%"+gbi.Description+"%")
	}
	result := query.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	getBookOutputs := make([]GetBookOutput, result.RowsAffected)
	for i, book := range books {
		getBookOutputs[i] = transformToDTO(book)
	}
	return getBookOutputs, nil
}

func transformToDTO(book model.Book) GetBookOutput {
	return GetBookOutput{
		ID:          book.ID.String(),
		Title:       book.Title,
		Description: book.Description,
	}
}
