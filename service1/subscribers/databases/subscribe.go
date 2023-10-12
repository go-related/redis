package databases

import (
	"context"
	"fmt"
	booksmodel "github.com/go-related/redis/service1/books/model"
	"github.com/go-related/redis/service1/subscribers/model"
)

func (s *subscribeDB) Subscribe(ctx context.Context, subscriberID uint, listOfBooks *[]booksmodel.Book, listOfAuthors *[]booksmodel.Author) (*model.Subscribe, error) {

	subscribeCurrentData, err := s.getSubscribeBySubscriberId(ctx, subscriberID)
	if err != nil {
		return nil, err
	}
	subscriber, err := s.GetSubscriberById(ctx, subscriberID)
	if err != nil {
		return nil, err
	}

	if subscribeCurrentData == nil {
		subscribeCurrentData = &model.Subscribe{
			ID:         uint(len(s.Subscribes)) + 1,
			Subscriber: *subscriber,
			IsDeleted:  false,
			Books:      listOfBooks,
			Authors:    listOfAuthors,
		}
		s.Subscribes = append(s.Subscribes, subscribeCurrentData)
	} else {
		subscribeCurrentData.Books = listOfBooks
		subscribeCurrentData.Authors = listOfAuthors
	}
	return subscribeCurrentData, nil
}

func (s *subscribeDB) DeleteSubscribe(ctx context.Context, Id uint) error {
	currentData, err := s.GetSubscribeById(ctx, Id)
	if err != nil {
		return err
	}
	currentData.IsDeleted = true
	return nil
}

func (s *subscribeDB) GetAllSubscribes(ctx context.Context) ([]model.Subscribe, error) {
	var result []model.Subscribe
	for _, dt := range s.Subscribes {
		if !dt.IsDeleted {
			result = append(result, *dt)
		}
	}
	return result, nil
}

func (s *subscribeDB) GetSubscribeById(ctx context.Context, Id uint) (*model.Subscribe, error) {
	for _, dt := range s.Subscribes {
		if dt.ID == Id {
			return dt, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}

func (s *subscribeDB) getSubscribeBySubscriberId(ctx context.Context, subscriberId uint) (*model.Subscribe, error) {
	for _, dt := range s.Subscribes {
		if dt.Subscriber.ID == subscriberId {
			return dt, nil
		}
	}
	return nil, nil // we are not returning an error here intentinally
}
