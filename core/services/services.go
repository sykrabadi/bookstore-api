package services

import (
	"bookstore-api/core/domain"
	"bookstore-api/core/ports"
)

type BookService struct {
	repo ports.BookRepository
}

func NewBookService (repo ports.BookRepository) *BookService{
	return &BookService{
		repo: repo,
	}
}

func (b *BookService) AddBook(book domain.Book) error{
	return b.repo.AddBook(book)
}
