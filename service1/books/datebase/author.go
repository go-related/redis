package datebase

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
)

func (b *booksDb) CreateAuthor(ctx context.Context, data model.Author) error {
	//TODO implement me
	panic("implement me")
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

func (b *booksDb) SearchAuthorsByName(ctx context.Context, title string) ([]model.Author, error) {
	//TODO implement me
	panic("implement me")
}
