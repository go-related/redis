package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-related/redis/redis"
	bookmodel "github.com/go-related/redis/service1/books/model"
	"github.com/go-related/redis/service1/database"
	"github.com/go-related/redis/service1/middleware"
	"github.com/go-related/redis/service1/subscribers/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// SubscriberHandler implements crud for handles GET /v1/api/genres
type SubscriberHandler struct {
	DB     database.BooksDB
	Engine *gin.Engine
	Redis  *redis.RedisService
}

func NewSubscriberHandler(db database.BooksDB, router *gin.Engine, rds *redis.RedisService) *SubscriberHandler {
	handler := &SubscriberHandler{
		Engine: router,
		DB:     db,
		Redis:  rds,
	}
	v1 := router.Group("/v1/api")
	// register subscribers
	v1.GET("/subscribers", handler.GetSubscribers)
	v1.GET("/subscribers/:id", handler.GetSubscriber)
	v1.PUT("/subscribers/:id", handler.UpdateSubscriber)
	v1.POST("/subscribers", handler.CreateSubscriber)
	v1.DELETE("/subscribers/:id", handler.DeleteSubscriber)

	// register subscribe
	v1.GET("/subscribes", handler.GetSubscribes)
	v1.GET("/subscribes/:id", handler.GetSubscribe)
	v1.POST("/subscribes", handler.CreateSubscribe)
	v1.DELETE("/subscribes/:id", handler.DeleteSubscribe)
	return handler
}

func (h *SubscriberHandler) GetSubscribers(c *gin.Context) {
	type QueryParameter struct {
		Limit  string `form:"limit,default=5" binding:"numeric"`
		Offset string `form:"offset,default=0" binding:"numeric"`
	}
	//TODO make uses of the pagination
	result, err := h.DB.GetAllSubscribers(context.Background())
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

func (h *SubscriberHandler) GetSubscriber(c *gin.Context) {
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
	result, err := h.DB.GetSubscriberById(context.Background(), uint(idValue))
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

func (h *SubscriberHandler) UpdateSubscriber(c *gin.Context) {
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
	type subscriber struct {
		Email    string `json:"email"`
		FullName string `json:"name"`
	}
	var input subscriber
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
	toUpdateData := model.Subscriber{
		ID:       uint(idValue),
		Email:    input.Email,
		FullName: input.FullName,
	}
	err = h.DB.UpdateSubscriber(context.Background(), toUpdateData)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusOK, toUpdateData)
}

func (h *SubscriberHandler) CreateSubscriber(c *gin.Context) {

	type subscriber struct {
		Email    string `json:"email"`
		FullName string `json:"name"`
	}
	var input subscriber
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
	toUpdateData := model.Subscriber{
		Email:    input.Email,
		FullName: input.FullName,
	}
	currentData, err := h.DB.CreateSubscriber(context.Background(), toUpdateData)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusCreated, currentData)
}

func (h *SubscriberHandler) DeleteSubscriber(c *gin.Context) {
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

	err = h.DB.DeleteSubscriber(context.Background(), uint(idValue))
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

func (h *SubscriberHandler) DeleteSubscribe(c *gin.Context) {
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

	err = h.DB.DeleteSubscribe(context.Background(), uint(idValue))
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

func (h *SubscriberHandler) GetSubscribes(c *gin.Context) {
	type QueryParameter struct {
		Limit  string `form:"limit,default=5" binding:"numeric"`
		Offset string `form:"offset,default=0" binding:"numeric"`
	}
	//TODO make uses of the pagination
	result, err := h.DB.GetAllSubscribes(context.Background())
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

func (h *SubscriberHandler) GetSubscribe(c *gin.Context) {
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
	result, err := h.DB.GetSubscribeById(context.Background(), uint(idValue))
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

func (h *SubscriberHandler) CreateSubscribe(c *gin.Context) {

	type subscribe struct {
		SubscriberID uint    `json:"subscriber_id"`
		Books        *[]uint `json:"books"`
		Authors      *[]uint `json:"authors"`
	}
	var input subscribe
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

	// TODO on here we need to load all books with one call
	books := make([]bookmodel.Book, 0)
	if input.Books != nil {
		for _, bookId := range *input.Books {
			bookData, err := h.DB.GetBookById(context.Background(), bookId)
			if err != nil {
				errorData := middleware.Response{
					StatusCode: http.StatusBadRequest,
					Err:        err,
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
				return
			}
			books = append(books, *bookData)
		}
	}
	authors := make([]bookmodel.Author, 0)
	if input.Authors != nil {
		for _, authorID := range *input.Authors {
			authorData, err := h.DB.GetAuthorById(context.Background(), authorID)
			if err != nil {
				errorData := middleware.Response{
					StatusCode: http.StatusBadRequest,
					Err:        err,
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, errorData)
				return
			}
			authors = append(authors, *authorData)
		}
	}

	currentData, err := h.DB.Subscribe(context.Background(), input.SubscriberID, &books, &authors)
	if err != nil {
		errorData := middleware.Response{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorData)
		return
	}
	c.IndentedJSON(http.StatusOK, currentData)
}
