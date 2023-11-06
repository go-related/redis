package databases

import (
	"context"
	"fmt"
	booksmodel "github.com/go-related/redis/service1/books/model"
	"github.com/go-related/redis/service1/subscribers/model"
)

type SubscribeDB interface {
	CreateSubscriber(ctx context.Context, data model.Subscriber) (model.Subscriber, error)
	UpdateSubscriber(ctx context.Context, data model.Subscriber) error
	DeleteSubscriber(ctx context.Context, Id uint) error
	GetAllSubscribers(ctx context.Context) ([]model.Subscriber, error)
	GetSubscriberById(ctx context.Context, Id uint) (*model.Subscriber, error)

	Subscribe(ctx context.Context, subscriberID uint, listOfBooks *[]booksmodel.Book, listOfAuthors *[]booksmodel.Author) (*model.Subscribe, error)
	DeleteSubscribe(ctx context.Context, Id uint) error
	GetAllSubscribes(ctx context.Context) ([]model.Subscribe, error)
	GetSubscribeById(ctx context.Context, Id uint) (*model.Subscribe, error)
	GetAuthorsSubscribers(ctx context.Context, listOfAuthors []booksmodel.Author) ([]model.Subscriber, error)
}

type subscribeDB struct {
	Subscribers []*model.Subscriber
	Subscribes  []*model.Subscribe
}

func NewSubscriberDB() SubscribeDB {
	return &subscribeDB{
		Subscribes:  []*model.Subscribe{},
		Subscribers: []*model.Subscriber{},
	}
}

func (s *subscribeDB) CreateSubscriber(ctx context.Context, data model.Subscriber) (model.Subscriber, error) {
	data.ID = uint(len(s.Subscribers) + 1)
	data.IsDeleted = false
	s.Subscribers = append(s.Subscribers, &data)
	return data, nil
}

func (s *subscribeDB) UpdateSubscriber(ctx context.Context, data model.Subscriber) error {
	currentData, err := s.GetSubscriberById(ctx, data.ID)
	if err != nil {
		return err
	}
	currentData.Email = data.Email
	currentData.FullName = data.FullName
	return nil
}

func (s *subscribeDB) DeleteSubscriber(ctx context.Context, Id uint) error {
	currentData, err := s.GetSubscriberById(ctx, Id)
	if err != nil {
		return err
	}
	currentData.IsDeleted = true
	return nil
}

func (s *subscribeDB) GetAllSubscribers(ctx context.Context) ([]model.Subscriber, error) {
	var result []model.Subscriber
	for _, dt := range s.Subscribers {
		if !dt.IsDeleted {
			result = append(result, *dt)
		}
	}
	return result, nil
}

func (s *subscribeDB) GetSubscriberById(ctx context.Context, Id uint) (*model.Subscriber, error) {
	for _, dt := range s.Subscribers {
		if dt.ID == Id {
			return dt, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}

func (s *subscribeDB) GetAuthorsSubscribers(ctx context.Context, listOfAuthors []booksmodel.Author) ([]model.Subscriber, error) {
	// helper to check if a subscriber is interested in our list of authors
	isSubscriberInterested := func(sub *model.Subscribe, listOfAuthors []booksmodel.Author) bool {
		if sub.Authors != nil {
			for _, potentialAuthor := range *sub.Authors {
				for _, interestedAuthor := range listOfAuthors {
					if potentialAuthor.ID == interestedAuthor.ID {
						return true
					}
				}
			}
		}

		return false
	}
	// helper to not register duplicated data on our potential result
	isSubscriberAlreadyRegistered := func(sub model.Subscriber, result []model.Subscriber) bool {
		for _, currentSub := range result {
			if sub.ID == currentSub.ID {
				return true
			}
		}
		return false
	}

	var result []model.Subscriber
	if len(listOfAuthors) > 0 {
		for _, dt := range s.Subscribes {
			if !dt.IsDeleted && isSubscriberInterested(dt, listOfAuthors) && !isSubscriberAlreadyRegistered(dt.Subscriber, result) {
				result = append(result, dt.Subscriber)
			}
		}
	}

	return result, nil

}
