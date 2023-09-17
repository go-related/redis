package model

type Book struct {
	ID        uint
	Title     string
	IsDeleted bool
	Authors   []Author
	Genres    []Genre
}

type Author struct {
	ID         uint
	IsDeleted  bool
	PublicName string
}

type Genre struct {
	ID        uint
	Name      string
	IsDeleted bool
}
