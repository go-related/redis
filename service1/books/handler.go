package books

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-related/redis/service1/books/datebase"
	"github.com/go-related/redis/service1/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func NewHandler(bookDb datebase.BooksDB) *Handler {
	return &Handler{
		bookDb: bookDb,
	}
}

type Handler struct {
	bookDb datebase.BooksDB
}

// GetBooks handles GET /v1/api/books
func (h *Handler) GetGenres(c *gin.Context) {
	type QueryParameter struct {
		Limit  string `form:"limit,default=5" binding:"numeric"`
		Offset string `form:"offset,default=0" binding:"numeric"`
	}
	//TODO make uses of the pagination
	result, err := h.bookDb.GetAllGenres(context.Background())
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
	result, err := h.bookDb.GetGenresById(context.Background(), uint(idValue))
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
