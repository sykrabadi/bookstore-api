package repository

import (
	"bookstore-api/core/domain"
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

type BookFirestoreRepository struct {
	client firestore.Client
}

func NewBookFirestoreRepository(ctx context.Context) *BookFirestoreRepository{
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("[NewBookFirestoreRepository]error load .env file with error %v \n", err)
	}

	projectId := os.Getenv("GCP_PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Printf("NewBookFirestoreRepository error with %v \n", err)
	}

	return &BookFirestoreRepository{
		client: *client,
	}
}

func (b *BookFirestoreRepository) AddBook(book domain.Book) error {
	ctx := context.Background()
	_, _, err := b.client.Collection("books").Add(ctx, book)

	if err != nil {
		log.Printf("error writing operation to books collection with error %v \n", err)
		return err
	}

	return nil
}
