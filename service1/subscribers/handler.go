package subscribers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-related/redis/service1/books/datebase"
)

// Handler implements crud for handles GET /v1/api/genres
type Handler struct {
	BookDb datebase.BooksDB
	Engine *gin.Engine
}
