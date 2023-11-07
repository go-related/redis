package database

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
	smodel "github.com/go-related/redis/service1/subscribers/model"
	"github.com/sirupsen/logrus"
)

func (b *booksDb) Subscribe(ctx context.Context, subscriberID uint, listOfBooks []model.Book, listOfAuthors []model.Author) (*smodel.Subscribe, error) {

	data := smodel.Subscribe{
		SubscriberID: subscriberID,
		Books:        listOfBooks,
		Authors:      listOfAuthors,
	}
	result := b.Db.Model(&smodel.Subscribe{}).Create(&data)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't create the subscribe")
	}
	return &data, result.Error
}

func (b *booksDb) DeleteSubscribe(ctx context.Context, Id uint) error {
	currentData, err := b.GetSubscribeById(ctx, Id)
	if err != nil {
		return err
	}
	result := b.Db.Delete(&currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", Id).Error("Error deleting subscribe")
	}
	return nil
}

func (b *booksDb) GetAllSubscribes(ctx context.Context) ([]*smodel.Subscribe, error) {
	var outputList []*smodel.Subscribe
	result := b.Db.Model(&smodel.Subscribe{}).Find(&outputList)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load subscribes")
	}
	return outputList, result.Error
}

func (b *booksDb) GetSubscribeById(ctx context.Context, Id uint) (*smodel.Subscribe, error) {
	var output smodel.Subscribe
	result := b.Db.Model(&smodel.Subscribe{}).First(&output, Id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load subscribe")
	}
	return &output, result.Error
}

func (b *booksDb) GetAuthorsSubscribers(ctx context.Context, listOfAuthors []model.Author) ([]*smodel.Subscriber, error) {
	var outputList []*smodel.Subscriber
	if len(listOfAuthors) == 0 {
		return outputList, nil
	}
	var authors []uint
	for _, author := range listOfAuthors {
		authors = append(authors, author.ID)
	}

	result := b.Db.Model(&smodel.Subscriber{}).Where("author_id in ? ", authors).Find(&outputList)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load subscribers for the authors")
	}
	return outputList, result.Error
}
