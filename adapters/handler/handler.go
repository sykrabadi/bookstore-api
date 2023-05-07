package handler

import (
	"bookstore-api/core/domain"
	"bookstore-api/core/services"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func (h *BookHTTPHandler) ViewBookByISBNHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		log.Println("[BookHTTPHandler.ViewBookByISBNHandler] invalid method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path, err := url.Parse(r.URL.Path)
	if err != nil {
		log.Printf("[BookHTTPHandler.ViewBookByISBNHandler] error parsing path with error %v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ISBN := strings.Replace(path.Path, "/books/", "", -1)
	data, err := h.bookService.ViewBookByISBN(ISBN)

	if err != nil {
		log.Printf("[BookHTTPHandler.ViewBookByISBNHandler] error occured when fetching data from FireStore with error %v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(data)
}

func (h *BookHTTPHandler) ViewBooks(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		log.Println("[BookHTTPHandler.ViewBookByISBNHandler] invalid method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := h.bookService.ViewBooks()

	if err != nil {
		log.Printf("[BookHTTPHandler.ViewBooks] error occured when fetching data from FireStore with error %v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(data)
}
