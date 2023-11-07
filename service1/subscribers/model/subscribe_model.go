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
	Subscriber Subscriber     `gorm:"many2many:subscribe_subscriber;"`
	Books      []model.Book   `gorm:"many2many:subscribe_book;"`
	Authors    []model.Author `gorm:"many2many:subscribe_author;"`
}
