package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-related/redis/service1/books/model"
	"github.com/go-related/redis/service1/middleware"
	"github.com/go-related/redis/settings"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) GetBooks(c *gin.Context) {
	type QueryParameter struct {
		Limit  string `form:"limit,default=5" binding:"numeric"`
		Offset string `form:"offset,default=0" binding:"numeric"`
	}
	//TODO make uses of the pagination
	result, err := h.BookDb.GetAllBooks(context.Background())
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) GetBook(c *gin.Context) {
	id := c.Params.ByName("id")

	idValue, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.WithError(err).Error("error converting id to int")
		errorData := middleware.Response{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
		return
	}
	result, err := h.BookDb.GetBookById(context.Background(), uint(idValue))
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")

	idValue, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.WithError(err).Error("error converting id to int")
		errorData := middleware.Response{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
		return
	}
	type Book struct {
		Title   string `json:"title"`
		Authors []uint `json:"authors"`
		Genres  []uint `json:"genres"`
	}
	var input Book
	err = c.BindJSON(&input)
	if err != nil {
		log.WithError(err).Error("error binding to json")
		errorData := middleware.Response{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
		return
	}
	//TODO check if the list of Authors and genres is valid
	// implement is valid for the Model
	bookData := model.Book{
		Title:   input.Title,
		Authors: []model.Author{},
		Genres:  []model.Genre{},
	}
	bookData.ID = uint(idValue)
	for _, dt := range input.Authors {
		authorDt, err := h.BookDb.GetAuthorById(context.Background(), dt)
		if err != nil {
			log.WithError(err).Error("invalid author Id")
			errorData := middleware.Response{
				StatusCode: http.StatusBadRequest,
				Err:        err,
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
			return
		}
		bookData.Authors = append(bookData.Authors, *authorDt)
	}
	for _, dt := range input.Genres {
		genreDt, err := h.BookDb.GetGenresById(context.Background(), dt)
		if err != nil {
			log.WithError(err).Error("invalid genre Id")
			errorData := middleware.Response{
				StatusCode: http.StatusBadRequest,
				Err:        err,
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
			return
		}
		bookData.Genres = append(bookData.Genres, *genreDt)
	}
	err = h.BookDb.UpdateBook(context.Background(), bookData)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusOK, bookData)
}

func (h *Handler) CreateBook(c *gin.Context) {
	type Book struct {
		Title   string `json:"title"`
		Authors []uint `json:"authors"`
		Genres  []uint `json:"genres"`
	}
	var input Book
	err := c.BindJSON(&input)
	if err != nil {
		log.WithError(err).Error("error binding to json")
		errorData := middleware.Response{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
		return
	}
	//TODO check if the list of Authors and genres is valid
	// implement is valid for the Model
	bookData := model.Book{
		Title:   input.Title,
		Authors: []model.Author{},
		Genres:  []model.Genre{},
	}
	for _, dt := range input.Authors {
		authorDt, err := h.BookDb.GetAuthorById(context.Background(), dt)
		if err != nil {
			log.WithError(err).Error("invalid author Id")
			errorData := middleware.Response{
				StatusCode: http.StatusBadRequest,
				Err:        err,
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
			return
		}
		bookData.Authors = append(bookData.Authors, *authorDt)
	}
	for _, dt := range input.Genres {
		genreDt, err := h.BookDb.GetGenresById(context.Background(), dt)
		if err != nil {
			log.WithError(err).Error("invalid genre Id")
			errorData := middleware.Response{
				StatusCode: http.StatusBadRequest,
				Err:        err,
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
			return
		}
		bookData.Genres = append(bookData.Genres, *genreDt)
	}
	data, err := h.BookDb.CreateBook(context.Background(), bookData)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	go h.publishNewBookToChannel(data)
	c.IndentedJSON(http.StatusCreated, data)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")
	idValue, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.WithError(err).Error("error converting id to int")
		errorData := middleware.Response{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
		return
	}

	err = h.BookDb.DeleteBook(context.Background(), uint(idValue))
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{
		"message": "Resource deleted successfully",
	})
}

func (h *Handler) publishNewBookToChannel(book model.Book) {
	payload, err := json.Marshal(book)
	if err != nil {
		log.WithError(err).Error("error converting model to string")
		return
	}
	err = h.Redis.PublishToChannel(context.Background(), settings.ApplicationConfiguration.Service1.NewBookChannelName, payload)
	if err != nil {
		log.WithError(err).Error("error publishing NewBook to channel")
	}
}
