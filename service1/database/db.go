package database

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
	smodel "github.com/go-related/redis/service1/subscribers/model"
	"github.com/go-related/redis/settings"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BooksDB interface {
	CreateAuthor(ctx context.Context, data model.Author) (model.Author, error)
	UpdateAuthor(ctx context.Context, data model.Author) error
	DeleteAuthor(ctx context.Context, Id uint) error
	GetAllAuthors(ctx context.Context) ([]*model.Author, error)
	GetAuthorById(ctx context.Context, Id uint) (*model.Author, error)

	CreateGenre(ctx context.Context, data model.Genre) (model.Genre, error)
	UpdateGenre(ctx context.Context, data model.Genre) error
	DeleteGenre(ctx context.Context, Id uint) error
	GetAllGenres(ctx context.Context) ([]*model.Genre, error)
	GetGenresById(ctx context.Context, Id uint) (*model.Genre, error)

	CreateBook(ctx context.Context, data model.Book) (model.Book, error)
	UpdateBook(ctx context.Context, data model.Book) error
	DeleteBook(ctx context.Context, Id uint) error
	GetAllBooks(ctx context.Context) ([]*model.Book, error)
	GetBookById(ctx context.Context, Id uint) (*model.Book, error)

	CreateSubscriber(ctx context.Context, data smodel.Subscriber) (smodel.Subscriber, error)
	UpdateSubscriber(ctx context.Context, data smodel.Subscriber) error
	DeleteSubscriber(ctx context.Context, Id uint) error
	GetAllSubscribers(ctx context.Context) ([]*smodel.Subscriber, error)
	GetSubscriberById(ctx context.Context, Id uint) (*smodel.Subscriber, error)

	Subscribe(ctx context.Context, subscriberID uint, listOfBooks *[]model.Book, listOfAuthors *[]model.Author) (*smodel.Subscribe, error)
	DeleteSubscribe(ctx context.Context, Id uint) error
	GetAllSubscribes(ctx context.Context) ([]*smodel.Subscribe, error)
	GetSubscribeById(ctx context.Context, Id uint) (*smodel.Subscribe, error)
	GetAuthorsSubscribers(ctx context.Context, listOfAuthors []*model.Author) ([]*smodel.Subscriber, error)
}

type booksDb struct {
	Db *gorm.DB
}

func NewBooks() (BooksDB, error) {
	result := booksDb{}
	db, err := gorm.Open(postgres.Open(settings.ApplicationConfiguration.Service1.DbConnectionString), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Error("error connecting to db")
		return &result, err
	}
	err = db.AutoMigrate(&model.Book{}, &model.Author{}, &model.Genre{}, &smodel.Subscriber{}, &smodel.Subscribe{})
	if err != nil {
		logrus.WithError(err).Error("couldn't migrate the db")
		return &result, err
	}
	result.Db = db
	return &result, nil
}
