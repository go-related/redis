package database

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
	"github.com/sirupsen/logrus"
)

func (b *booksDb) CreateBook(ctx context.Context, data model.Book) (model.Book, error) {
	result := b.Db.Model(&model.Book{}).Create(&data)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't create the book")
	}
	return data, result.Error
}

func (b *booksDb) UpdateBook(ctx context.Context, data model.Book) error {
	currentData, err := b.GetBookById(ctx, data.ID)
	if err != nil {
		// we dont log here since will be logged on getbyID
		return err
	}
	currentData.Title = data.Title
	currentData.Authors = data.Authors
	currentData.Genres = data.Genres
	result := b.Db.Save(currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", data.ID).Error("Error updating book")
	}
	return result.Error
}

func (b *booksDb) DeleteBook(ctx context.Context, Id uint) error {
	currentData, err := b.GetBookById(ctx, Id)
	if err != nil {
		// we dont log here since will be logged on getbyID
		return err
	}
	result := b.Db.Delete(&currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", Id).Error("Error deleting book")
	}
	return nil
}

func (b *booksDb) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	var outputList []*model.Book
	result := b.Db.Model(&model.Book{}).Find(&outputList)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load books")
	}
	return outputList, result.Error
}

func (b *booksDb) GetBookById(ctx context.Context, Id uint) (*model.Book, error) {
	var output model.Book
	result := b.Db.Model(&model.Book{}).Preload("Authors").Preload("Genres").First(&output, Id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load book")
	}
	return &output, result.Error
}
