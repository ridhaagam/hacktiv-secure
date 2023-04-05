// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	bookService "secure/challenge-2/application/usecases/book"
	bookRepository "secure/challenge-2/infrastructure/repository/postgres/book"
	bookController "secure/challenge-2/infrastructure/restapi/controllers/book"

	"gorm.io/gorm"
)

// BookAdapter is a function that returns a book controller
func BookAdapter(db *gorm.DB) *bookController.Controller {
	mRepository := bookRepository.Repository{DB: db}
	service := bookService.Service{BookRepository: mRepository}
	return &bookController.Controller{BookService: service}
}
