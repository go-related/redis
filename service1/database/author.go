package database

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
	"github.com/sirupsen/logrus"
)

func (b *booksDb) CreateAuthor(ctx context.Context, data model.Author) (model.Author, error) {
	result := b.Db.Model(&model.Author{}).Create(&data)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't create the author")
	}
	return data, result.Error
}

func (b *booksDb) UpdateAuthor(ctx context.Context, data model.Author) error {
	currentData, err := b.GetAuthorById(ctx, data.ID)
	if err != nil {
		// we dont log here since will be logged on getbyID
		return err
	}
	currentData.PublicName = data.PublicName
	result := b.Db.Save(currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", data.ID).Error("Error updating author")
	}
	return result.Error
}

func (b *booksDb) DeleteAuthor(ctx context.Context, Id uint) error {
	currentData, err := b.GetAuthorById(ctx, Id)
	if err != nil {
		// we dont log here since will be logged on getbyID
		return err
	}
	result := b.Db.Delete(&currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", Id).Error("Error deleting author")
	}
	return nil
}

func (b *booksDb) GetAllAuthors(ctx context.Context) ([]*model.Author, error) {
	var outputList []*model.Author
	result := b.Db.Model(&model.Author{}).Find(&outputList)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load authors")
	}
	return outputList, result.Error
}

func (b *booksDb) GetAuthorById(ctx context.Context, Id uint) (*model.Author, error) {
	var output model.Author
	result := b.Db.Model(&model.Author{}).First(&output, Id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load author")
	}
	return &output, result.Error
}
