package main

import (
	"bookstore-api/adapters/handler"
	"bookstore-api/adapters/repository"
	"bookstore-api/core/services"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func InitHttpServer(bookService services.BookService) {
	mux := http.NewServeMux()
	bookHandler := handler.NewBookHTTPHandler(bookService)
	mux.HandleFunc("/books", bookHandler.AddBookHandler)

	port := 8000
	server := http.Server{
		Addr: fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Printf("[InitHttpServer] fail to initialize http server with error %v \n", err)
	}

	log.Printf("listening at port %v \n", port)
}

func main() {
	ctx := context.Background()
	store := repository.NewBookFirestoreRepository(ctx)
	bookService := services.NewBookService(store)
	
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go InitHttpServer(*bookService)
	<-done
}
