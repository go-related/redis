package datebase

import (
	"context"
	"fmt"
	"github.com/go-related/redis/service1/books/model"
	"strings"
)

func (b *booksDb) CreateAuthor(_ context.Context, data model.Author) (model.Author, error) {
	data.ID = uint(len(b.Authors) + 1)
	data.IsDeleted = false
	b.Authors = append(b.Authors, &data)
	return data, nil
}

func (b *booksDb) UpdateAuthor(ctx context.Context, data model.Author) error {
	currentData, err := b.GetAuthorById(ctx, data.ID)
	if err != nil {
		return err
	}
	currentData.PublicName = data.PublicName
	return nil
}

func (b *booksDb) DeleteAuthor(ctx context.Context, Id uint) error {
	currentData, err := b.GetAuthorById(ctx, Id)
	if err != nil {
		return err
	}
	currentData.IsDeleted = true
	return nil
}

func (b *booksDb) GetAllAuthors(_ context.Context) ([]model.Author, error) {
	var result []model.Author
	for _, dt := range b.Authors {
		if !dt.IsDeleted {
			result = append(result, *dt)
		}
	}
	return result, nil
}

func (b *booksDb) SearchAuthorsByName(_ context.Context, title string) ([]model.Author, error) {
	var result []model.Author
	for _, dt := range b.Authors {
		if !dt.IsDeleted && strings.Contains(dt.PublicName, title) {
			result = append(result, *dt)
		}
	}
	return result, nil
}

func (b *booksDb) GetAuthorById(_ context.Context, id uint) (*model.Author, error) {
	for _, dt := range b.Authors {
		if dt.ID == id {
			return dt, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}
