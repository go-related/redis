package database

import (
	"context"
	smodel "github.com/go-related/redis/service1/subscribers/model"
)

func (b *booksDb) CreateSubscriber(ctx context.Context, data smodel.Subscriber) (smodel.Subscriber, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) UpdateSubscriber(ctx context.Context, data smodel.Subscriber) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) DeleteSubscriber(ctx context.Context, Id uint) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetAllSubscribers(ctx context.Context) ([]smodel.Subscriber, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetSubscriberById(ctx context.Context, Id uint) (*smodel.Subscriber, error) {
	//TODO implement me
	panic("implement me")
}
