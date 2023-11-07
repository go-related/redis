package database

import (
	"context"
	smodel "github.com/go-related/redis/service1/subscribers/model"
	"github.com/sirupsen/logrus"
)

func (b *booksDb) CreateSubscriber(ctx context.Context, data smodel.Subscriber) (smodel.Subscriber, error) {
	result := b.Db.Model(&smodel.Subscriber{}).Create(&data)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't create the subscriber")
	}
	return data, result.Error
}

func (b *booksDb) UpdateSubscriber(ctx context.Context, data smodel.Subscriber) error {
	currentData, err := b.GetSubscriberById(ctx, data.ID)
	if err != nil {
		return err
	}
	currentData.FullName = data.FullName
	currentData.Email = data.Email
	result := b.Db.Save(currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", data.ID).Error("Error updating subscriber")
	}
	return result.Error
}

func (b *booksDb) DeleteSubscriber(ctx context.Context, Id uint) error {
	currentData, err := b.GetSubscriberById(ctx, Id)
	if err != nil {
		return err
	}
	result := b.Db.Delete(&currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", Id).Error("Error deleting subscriber")
	}
	return nil
}

func (b *booksDb) GetAllSubscribers(ctx context.Context) ([]*smodel.Subscriber, error) {
	var outputList []*smodel.Subscriber
	result := b.Db.Model(&smodel.Subscriber{}).Find(&outputList)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load subscribers")
	}
	return outputList, result.Error
}

func (b *booksDb) GetSubscriberById(ctx context.Context, Id uint) (*smodel.Subscriber, error) {
	var output smodel.Subscriber
	result := b.Db.Model(&smodel.Subscriber{}).First(&output, Id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load subscriber")
	}
	return &output, result.Error
}
