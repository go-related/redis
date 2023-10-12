package databases

import (
	"context"
	"github.com/go-related/redis/service1/subscribers/model"
)

type SubscribeDB interface {
	CreateSubscriber(ctx context.Context, data model.Subscriber) (model.Subscriber, error)
	UpdateSubscriber(ctx context.Context, data model.Subscriber) error
	DeleteSubscriber(ctx context.Context, Id uint) error
	GetAllSubscribers(ctx context.Context) ([]model.Subscriber, error)
	GetSubscriberById(ctx context.Context, Id uint) (model.Subscriber, error)

	Subscribe(ctx context.Context, data model.Subscribe) (model.Subscribe, error)
	DeleteAllSubscribes(ctx context.Context, Id uint) error
	GetAllSubscribes(ctx context.Context) ([]model.Subscribe, error)
	GetSubscribeById(ctx context.Context, Id uint) (model.Subscribe, error)
}

type subscribeDB struct {
	Subscribers []model.Subscriber
	Subscribes  []model.Subscribe
}

func NewSubscriberDB() SubscribeDB {
	return &subscribeDB{
		Subscribes:  []model.Subscribe{},
		Subscribers: []model.Subscriber{},
	}
}

func (s *subscribeDB) CreateSubscriber(ctx context.Context, data model.Subscriber) (model.Subscriber, error) {
	//TODO implement me
	panic("implement me")
}

func (s *subscribeDB) UpdateSubscriber(ctx context.Context, data model.Subscriber) error {
	//TODO implement me
	panic("implement me")
}

func (s *subscribeDB) DeleteSubscriber(ctx context.Context, Id uint) error {
	//TODO implement me
	panic("implement me")
}

func (s *subscribeDB) GetAllSubscribers(ctx context.Context) ([]model.Subscriber, error) {
	//TODO implement me
	panic("implement me")
}

func (s *subscribeDB) GetSubscriberById(ctx context.Context, Id uint) (model.Subscriber, error) {
	//TODO implement me
	panic("implement me")
}
