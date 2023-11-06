package database

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
	smodel "github.com/go-related/redis/service1/subscribers/model"
)

func (b *booksDb) Subscribe(ctx context.Context, subscriberID uint, listOfBooks *[]model.Book, listOfAuthors *[]model.Author) (*smodel.Subscribe, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) DeleteSubscribe(ctx context.Context, Id uint) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetAllSubscribes(ctx context.Context) ([]smodel.Subscribe, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetSubscribeById(ctx context.Context, Id uint) (*smodel.Subscribe, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetAuthorsSubscribers(ctx context.Context, listOfAuthors []model.Author) ([]smodel.Subscriber, error) {
	//TODO implement me
	panic("implement me")
}
