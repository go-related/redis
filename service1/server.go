package service1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-related/redis/redis"
	"github.com/go-related/redis/service1/books/handler"
	"github.com/go-related/redis/service1/database"
	handler2 "github.com/go-related/redis/service1/subscribers/handler"
	"github.com/go-related/redis/settings"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	appSettings settings.Service1
}

func InitService(appSettings settings.Service1) (*Service, error) {
	router := gin.Default()
	booksDb, err := database.NewBooks()
	if err != nil {
		log.WithError(err).Fatal("couldn't connect to db")
	}
	redisService := redis2.New()
	handler.NewHandler(booksDb, router, redisService)
	handler2.NewSubscriberHandler(booksDb, router, redisService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	err = router.Run(fmt.Sprintf(":%s", appSettings.Port))
	if err != nil {
		log.WithError(err).Errorf("Setting up service failed.")
		return nil, err
	}
	log.Infof("Application '%s' is running on port:%s", appSettings.Name, appSettings.Port)
	return &Service{appSettings: appSettings}, nil
}
