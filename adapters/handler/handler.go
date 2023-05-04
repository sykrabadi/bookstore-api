package handler

import (
	"bookstore-api/core/domain"
	"bookstore-api/core/services"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type BookHTTPHandler struct {
	bookService services.BookService
}

func NewBookHTTPHandler(bookService services.BookService) *BookHTTPHandler {
	return &BookHTTPHandler{
		bookService: bookService,
	}
}

func (h *BookHTTPHandler) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("[BookHTTPHandler.AddBookHandler] invalid method")
		return
	}
	payload := domain.Book{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[BookHTTPHandler.AddBookHandler] error reading request body with error %v \n", err)
		return
	}

	err = json.Unmarshal(body, &payload)

	if err != nil {
		log.Printf("[BookHTTPHandler.AddBookHandler] error unmarshaling request body with error %v \n", err)
		return
	}

	err = h.bookService.AddBook(payload)
	if err != nil {
		log.Printf("[BookHTTPHandler.AddBookHandler] error add book with error %v \n", err)
	}
}
