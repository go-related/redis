package model

import (
	"github.com/go-related/redis/service1/books/model"
	"gorm.io/gorm"
)

type Subscriber struct {
	gorm.Model
	Email    string
	FullName string
}

type Subscribe struct {
	gorm.Model
	Subscriber Subscriber
	Books      *[]model.Book
	Authors    *[]model.Author
}
