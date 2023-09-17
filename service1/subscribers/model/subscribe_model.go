package model

import "github.com/go-related/redis/service1/books/model"

type Subscriber struct {
	ID        uint
	Email     string
	FullName  string
	IsDeleted bool
}

type Subscribe struct {
	ID         uint
	IsDeleted  bool
	Subscriber Subscriber
	Genres     *[]model.Genre
	Authors    *[]model.Author
}
