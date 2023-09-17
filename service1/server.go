package service1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-related/redis/service1/books"
	"github.com/go-related/redis/service1/books/datebase"
	"github.com/go-related/redis/settings"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	appSettings settings.Service1
}

func InitService(appSettings settings.Service1) (*Service, error) {
	router := gin.Default()
	booksDb := datebase.NewBooksDB()
	booksHandler := books.NewHandler(booksDb)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.GET("/v1/api/genres", booksHandler.GetGenres)
	router.GET("/v1/api/genres/:id", booksHandler.GetGenre)
	err := router.Run(fmt.Sprintf(":%s", appSettings.Port))
	if err != nil {
		log.WithError(err).Errorf("Setting up service failed.")
		return nil, err
	}
	log.Infof("Application '%s' is running on port:%s", appSettings.Name, appSettings.Port)
	return &Service{appSettings: appSettings}, nil
}
