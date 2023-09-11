package main

import (
	"fmt"
	"log"

	// "github.com/oklog/ulid/v2"
	"github.com/samirreza/library/internal/model"
	"github.com/samirreza/library/internal/query"

	// "github.com/samirreza/library/internal/repository"
	// "github.com/samirreza/library/internal/service"
	// "github.com/samirreza/library/internal/service/DTO"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://postgres:root@localhost:5432/library"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&model.Book{})

	// bookRepository := repository.NewBookRepository(*db)

	// createBookService := service.NewCreateBookService(bookRepository)
	// id := ulid.Make().String()
	// createBookDTO := DTO.CreateBookDTO{
	// 	ID:          id,
	// 	Title:       "My First Book",
	// 	Description: "description for my first book",
	// }
	// createBookService.Create(createBookDTO)

	// updateBookService := service.NewUpdateBookService(bookRepository)
	// ubd := DTO.UpdateBookDTO{
	// 	Title: "My First Updated Book",
	// 	Description: "updated description for my first book",
	// }
	// updateBookService.Update(id, ubd)

	// getBookService := service.NewGetBookService(bookRepository)
	// book, _ := getBookService.Get(id)
	// fmt.Printf("%+v\n", book)

	// deletebookService := service.NewDeleteBookService(bookRepository)
	// deletebookService.Delete(book.ID)

	getBooksQuery := query.NewGetBooksQuery(*db)
	getBooksInput := query.GetBooksInput{
		PaginationInput: query.PaginationInput{
			Page:     2,
			PageSize: 1,
		},
		Title: "First",
		// Description: "second",
	}
	result, _ := getBooksQuery.Get(getBooksInput)
	fmt.Printf("%+v", result)
}
