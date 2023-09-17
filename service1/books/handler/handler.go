package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-related/redis/service1/books/datebase"
)

func NewHandler(bookDb datebase.BooksDB, router *gin.Engine) *Handler {
	handler := &Handler{
		BookDb: bookDb,
		Engine: router,
	}
	v1 := router.Group("/v1/api")
	// register genres
	v1.GET("/genres", handler.GetGenres)
	v1.GET("/genres/:id", handler.GetGenre)
	v1.PUT("/genres/:id", handler.UpdateGenre)
	v1.POST("/genres", handler.CreateGenre)
	v1.DELETE("/genres/:id", handler.DeleteGenre)

	// register authors
	v1.GET("/authors", handler.GetAuthors)
	v1.GET("/authors/:id", handler.GetAuthor)
	//v1.GET("/authors/:name", handler.GetAuthorsByName)
	v1.PUT("/authors/:id", handler.UpdateAuthor)
	v1.POST("/authors", handler.CreateAuthor)
	v1.DELETE("/authors/:id", handler.DeleteAuthor)

	//register books
	v1.GET("/books", handler.GetBooks)
	v1.GET("/books/:id", handler.GetBook)
	//v1.GET("/books/:title", handler.GetBooksByTitle)
	v1.PUT("/books/:id", handler.UpdateBook)
	v1.POST("/books", handler.CreateBook)
	v1.DELETE("/books/:id", handler.DeleteBook)

	return handler
}

// Handler implements crud for handles GET /v1/api/genres
type Handler struct {
	BookDb datebase.BooksDB
	Engine *gin.Engine
}
