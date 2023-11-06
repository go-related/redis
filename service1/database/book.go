package database

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
)

func (b *booksDb) CreateBook(ctx context.Context, data model.Book) (model.Book, error) {
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
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetBookById(ctx context.Context, Id uint) (*model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) SearchBooksByTitle(ctx context.Context, title string) ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}
