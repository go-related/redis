package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-related/redis/service1/books/model"
	"github.com/go-related/redis/service1/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

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
