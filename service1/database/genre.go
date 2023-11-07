package database

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
	"github.com/sirupsen/logrus"
)

func (b *booksDb) CreateGenre(ctx context.Context, data model.Genre) (model.Genre, error) {
	result := b.Db.Model(&model.Author{}).Create(&data)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't create the genre")
	}
	return data, result.Error
}

func (b *booksDb) UpdateGenre(ctx context.Context, data model.Genre) error {
	currentData, err := b.GetGenresById(ctx, data.ID)
	if err != nil {
		return err
	}
	currentData.Name = data.Name
	result := b.Db.Save(currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", data.ID).Error("Error updating genre")
	}
	return result.Error
}

func (b *booksDb) DeleteGenre(ctx context.Context, Id uint) error {
	currentData, err := b.GetGenresById(ctx, Id)
	if err != nil {
		return err
	}
	result := b.Db.Delete(&currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", Id).Error("Error deleting genre")
	}
	return nil
}

func (b *booksDb) GetAllGenres(ctx context.Context) ([]*model.Genre, error) {
	var outputList []*model.Genre
	result := b.Db.Model(&model.Genre{}).Find(&outputList)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load genres")
	}
	return outputList, result.Error
}

func (b *booksDb) GetGenresById(ctx context.Context, Id uint) (*model.Genre, error) {
	var output model.Genre
	result := b.Db.Model(&model.Genre{}).First(&output, Id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load genre")
	}
	return &output, result.Error
}
