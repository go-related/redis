package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title   string
	Authors []Author
	Genres  []Genre
}

type Author struct {
	gorm.Model
	PublicName string
}

type Genre struct {
	gorm.Model
	Name string
}
