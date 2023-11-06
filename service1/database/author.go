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
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) DeleteAuthor(ctx context.Context, Id uint) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetAllAuthors(ctx context.Context) ([]model.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetAuthorById(ctx context.Context, Id uint) (*model.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) SearchAuthorsByName(ctx context.Context, title string) ([]model.Author, error) {
	//TODO implement me
	panic("implement me")
}
