package datebase

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
)

type BooksDB interface {
	CreateAuthor(ctx context.Context, data model.Author) error
	UpdateAuthor(ctx context.Context, data model.Author) error
	DeleteAuthor(ctx context.Context, Id uint) error
	GetAllAuthors(ctx context.Context) ([]model.Author, error)
	SearchAuthorsByName(ctx context.Context, title string) ([]model.Author, error)

	CreateGenre(ctx context.Context, data model.Genre) error
	UpdateGenre(ctx context.Context, data model.Genre) error
	DeleteGenre(ctx context.Context, Id uint) error
	GetAllGenres(ctx context.Context) ([]model.Genre, error)

	CreateBook(ctx context.Context, data model.Book) error
	UpdateBook(ctx context.Context, data model.Book) error
	DeleteBook(ctx context.Context, Id uint) error
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	SearchBooksByTitle(ctx context.Context, title string) ([]model.Book, error)
}

type booksDb struct {
	Authors []*model.Author
	Books   []*model.Book
	Genre   []*model.Genre
}

func NewBooksDB() BooksDB {
	initialGenres := []*model.Genre{
		{
			ID:   1,
			Name: "Fiction",
		},
		{
			ID:   2,
			Name: "Science",
		},
		{
			ID:   3,
			Name: "Science Fiction",
		},
		{
			ID:   4,
			Name: "Romance",
		},
	}

	return &booksDb{
		Genre:   initialGenres,
		Authors: []*model.Author{},
		Books:   []*model.Book{},
	}
}

func (b *booksDb) CreateBook(ctx context.Context, data model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) UpdateBook(ctx context.Context, data model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) DeleteBook(ctx context.Context, Id uint) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	var result []model.Book

	for _, data := range b.Books {
		result = append(result, *data)
	}
	return result, nil
}

func (b *booksDb) SearchBooksByTitle(ctx context.Context, title string) ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}
