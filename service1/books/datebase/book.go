package datebase

import (
	"context"
	"fmt"
	"github.com/go-related/redis/service1/books/model"
	"strings"
)

type BooksDB interface {
	CreateAuthor(ctx context.Context, data model.Author) (model.Author, error)
	UpdateAuthor(ctx context.Context, data model.Author) error
	DeleteAuthor(ctx context.Context, Id uint) error
	GetAllAuthors(ctx context.Context) ([]model.Author, error)
	GetAuthorById(ctx context.Context, Id uint) (*model.Author, error)
	SearchAuthorsByName(ctx context.Context, title string) ([]model.Author, error)

	CreateGenre(ctx context.Context, data model.Genre) (model.Genre, error)
	UpdateGenre(ctx context.Context, data model.Genre) error
	DeleteGenre(ctx context.Context, Id uint) error
	GetAllGenres(ctx context.Context) ([]model.Genre, error)
	GetGenresById(ctx context.Context, Id uint) (*model.Genre, error)

	CreateBook(ctx context.Context, data model.Book) (model.Book, error)
	UpdateBook(ctx context.Context, data model.Book) error
	DeleteBook(ctx context.Context, Id uint) error
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	GetBookById(ctx context.Context, Id uint) (*model.Book, error)
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

func (b *booksDb) CreateBook(ctx context.Context, data model.Book) (model.Book, error) {
	data.ID = uint(len(b.Books) + 1)
	data.IsDeleted = false
	b.Books = append(b.Books, &data)
	return data, nil
}

func (b *booksDb) UpdateBook(ctx context.Context, data model.Book) error {
	currentData, err := b.GetBookById(ctx, data.ID)
	if err != nil {
		return err
	}
	currentData.Title = data.Title
	currentData.Authors = data.Authors //this is not supposed to be called on a patch
	currentData.Genres = data.Genres
	return nil
}

func (b *booksDb) DeleteBook(ctx context.Context, Id uint) error {
	currentData, err := b.GetBookById(ctx, Id)
	if err != nil {
		return err
	}
	currentData.IsDeleted = true
	return nil
}

func (b *booksDb) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	var result []model.Book

	for _, data := range b.Books {
		if !data.IsDeleted {
			result = append(result, *data)
		}
	}
	return result, nil
}

func (b *booksDb) SearchBooksByTitle(ctx context.Context, title string) ([]model.Book, error) {
	var result []model.Book

	for _, data := range b.Books {
		if !data.IsDeleted && strings.Contains(data.Title, title) {
			result = append(result, *data)
		}
	}
	return result, nil
}

func (b *booksDb) GetBookById(ctx context.Context, Id uint) (*model.Book, error) {
	for _, dt := range b.Books {
		if dt.ID == Id {
			return dt, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}
