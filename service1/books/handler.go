package books

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-related/redis/service1/books/datebase"
	"github.com/go-related/redis/service1/books/model"
	"github.com/go-related/redis/service1/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	v1.GET("/authors/:name", handler.GetAuthorsByName)
	v1.PUT("/authors/:id", handler.UpdateAuthor)
	v1.POST("/authors", handler.CreateAuthor)
	v1.DELETE("/authors/:id", handler.DeleteAuthor)

	//register books
	v1.GET("/books", handler.GetBooks)
	v1.GET("/books/:id", handler.GetBook)
	v1.GET("/books/:title", handler.GetBooksByTitle)
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

func (h *Handler) GetGenres(c *gin.Context) {
	type QueryParameter struct {
		Limit  string `form:"limit,default=5" binding:"numeric"`
		Offset string `form:"offset,default=0" binding:"numeric"`
	}
	//TODO make uses of the pagination
	result, err := h.BookDb.GetAllGenres(context.Background())
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

func (h *Handler) GetGenre(c *gin.Context) {
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
	result, err := h.BookDb.GetGenresById(context.Background(), uint(idValue))
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

func (h *Handler) UpdateGenre(c *gin.Context) {
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
	type Genre struct {
		Name string `json:"name"`
	}
	var input Genre
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
	genreData := model.Genre{
		ID:   uint(idValue),
		Name: input.Name,
	}
	err = h.BookDb.UpdateGenre(context.Background(), genreData)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusOK, genreData)
}

func (h *Handler) CreateGenre(c *gin.Context) {
	type Genre struct {
		Name string `json:"name"`
	}
	var input Genre
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
	genreData := model.Genre{
		Name: input.Name,
	}
	data, err := h.BookDb.CreateGenre(context.Background(), genreData)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusCreated, data)
}

func (h *Handler) DeleteGenre(c *gin.Context) {
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

	err = h.BookDb.DeleteGenre(context.Background(), uint(idValue))
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

func (h *Handler) GetAuthors(c *gin.Context) {
	type QueryParameter struct {
		Limit  string `form:"limit,default=5" binding:"numeric"`
		Offset string `form:"offset,default=0" binding:"numeric"`
	}
	//TODO make uses of the pagination
	result, err := h.BookDb.GetAllAuthors(context.Background())
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

func (h *Handler) GetAuthor(c *gin.Context) {
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
	result, err := h.BookDb.GetAuthorById(context.Background(), uint(idValue))
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

func (h *Handler) GetAuthorsByName(c *gin.Context) {
	name := c.Params.ByName("name")
	result, err := h.BookDb.SearchAuthorsByName(context.Background(), name)
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

func (h *Handler) UpdateAuthor(c *gin.Context) {
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
	type Author struct {
		Name string `json:"name"`
	}
	var input Author
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
	authorData := model.Author{
		ID:         uint(idValue),
		PublicName: input.Name,
	}
	err = h.BookDb.UpdateAuthor(context.Background(), authorData)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusOK, authorData)
}

func (h *Handler) CreateAuthor(c *gin.Context) {
	type Author struct {
		Name string `json:"name"`
	}
	var input Author
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
	authorData := model.Author{
		PublicName: input.Name,
	}
	data, err := h.BookDb.CreateAuthor(context.Background(), authorData)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusCreated, data)
}

func (h *Handler) DeleteAuthor(c *gin.Context) {
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

	err = h.BookDb.DeleteAuthor(context.Background(), uint(idValue))
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

func (h *Handler) GetBooksByTitle(c *gin.Context) {
	title := c.Params.ByName("title")
	result, err := h.BookDb.SearchBooksByTitle(context.Background(), title)
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
		ID:      uint(idValue),
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
