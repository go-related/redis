package datebase

import (
	"context"
	"github.com/go-related/redis/service1/books/model"
)

func (b *booksDb) CreateGenre(ctx context.Context, data model.Genre) error {
	data.ID = uint(len(b.Genre) + 1)
	data.IsDeleted = false
	b.Genre = append(b.Genre, &data)
	return nil
}

func (b *booksDb) UpdateGenre(ctx context.Context, data model.Genre) error {
	currentData, err := b.GetById(ctx, data.ID)
	if err != nil {
		return err
	}
	currentData.Name = data.Name
	return nil
}

func (b *booksDb) DeleteGenre(ctx context.Context, Id uint) error {
	currentData, err := b.GetById(ctx, Id)
	if err != nil {
		return err
	}
	currentData.IsDeleted = true
	return nil
}

func (b *booksDb) GetAllGenres(ctx context.Context) ([]model.Genre, error) {
	var result []model.Genre
	for _, dt := range b.Genre {
		if !dt.IsDeleted {
			result = append(result, *dt)
		}
	}
	return result, nil
}

func (b *booksDb) GetById(ctx context.Context, id uint) (*model.Genre, error) {
	for _, dt := range b.Genre {
		if dt.ID == id {
			return dt, nil
		}
	}
	return nil, nil
}
