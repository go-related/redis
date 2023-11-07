package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title   string
	Authors []Author `gorm:"many2many:book_authors;"`
	Genres  []Genre  `gorm:"many2many:book_genres;"`
}

type Author struct {
	gorm.Model
	PublicName string
}

type Genre struct {
	gorm.Model
	Name string
}
