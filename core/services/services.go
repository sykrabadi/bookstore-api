package services

import (
	"bookstore-api/core/domain"
	"bookstore-api/core/ports"
	"log"
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

func (b *BookService) ViewBookByISBN(ISBN string) (*domain.Book, error){
	res, err := b.repo.ViewBookByISBN(ISBN)
	if err != nil {
		log.Printf("[BookService.ViewBookByISBN] error with error %v \n", err)
		return nil, err
	}

	return res, nil
}

func (b *BookService) ViewBooks() ([]*domain.Book, error){
	res, err := b.repo.ViewBooks()
	if err != nil {
		log.Printf("[BookService.ViewBooks] error with error %v \n", err)
		return nil, err
	}

	return res, nil
}
