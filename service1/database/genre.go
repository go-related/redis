package database

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
)

func (b *booksDb) CreateGenre(ctx context.Context, data model.Genre) (model.Genre, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) UpdateGenre(ctx context.Context, data model.Genre) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) DeleteGenre(ctx context.Context, Id uint) error {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetAllGenres(ctx context.Context) ([]model.Genre, error) {
	//TODO implement me
	panic("implement me")
}

func (b *booksDb) GetGenresById(ctx context.Context, Id uint) (*model.Genre, error) {
	//TODO implement me
	panic("implement me")
}
