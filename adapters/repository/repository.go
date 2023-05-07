package repository

import (
	"bookstore-api/core/domain"
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

type BookFirestoreRepository struct {
	client firestore.Client
	pubsub pubsub.Client
}

func NewBookFirestoreRepository(ctx context.Context) *BookFirestoreRepository {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("[NewBookFirestoreRepository]error load .env file with error %v \n", err)
	}

	projectId := os.Getenv("GCP_PROJECT_ID")
	log.Println(projectId)
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("[NewBookFirestoreRepository] failed to initialize firestore client with error %v \n", err)
		return nil
	}
	pubsubClient, err := pubsub.NewClient(context.Background(), projectId)
	if err != nil {
		log.Fatalf("[NewBookFirestoreRepository] failed to initialize pubsub client with error %v \n", err)
		return nil
	}

	return &BookFirestoreRepository{
		client: *client,
		pubsub: *pubsubClient,
	}
}

func (b *BookFirestoreRepository) AddBook(book domain.Book) error {
	ctx := context.Background()
	_, _, err := b.client.Collection("books").Add(ctx, book)

	if err != nil {
		log.Printf("error writing operation to books collection with error %v \n", err)
		return err
	}

	topic := b.pubsub.Topic("upload-image")
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("book added"),
	})
	_, err = res.Get(ctx)
	if err != nil {
		log.Printf("[BookFirestoreRepository.AddBook] fail to publish message to %v topic with error %v \n", topic, err)
		return err
	}
	return nil
}

func (b *BookFirestoreRepository) ViewBookByISBN(ISBN string) (*domain.Book, error){
	ctx := context.Background()
	dsnap, err := b.client.Collection("books").Doc(ISBN).Get(ctx)

	if err != nil {
		return nil, err
	}
	log.Println(dsnap.Data())

	return nil, nil
}

func (b *BookFirestoreRepository) ViewBooks() ([]*domain.Book, error){
	var resp []*domain.Book
	ctx := context.Background()
	dsnaps, err := b.client.Collection("books").Limit(5).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	for _, data := range dsnaps{
		resp = append(resp, )
		log.Println(data.Data()["Author"])
	}

	return nil, nil
}
