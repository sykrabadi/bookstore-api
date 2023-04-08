package ports

import "bookstore-api/core/domain"

type BookService interface {
	AddBook(book domain.Book) error
	// UpdateBook(book domain.Book) error
	// DeleteBook(book domain.Book) error
	// ViewBooks() ([]*domain.Book, error)
}

type BookRepository interface {
	AddBook(book domain.Book) error
	// UpdateBook(book domain.Book) error
	// DeleteBook(book domain.Book) error
	// ViewBooks() ([]*domain.Book, error)
}