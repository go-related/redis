package service2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-related/redis/redis"
	"github.com/go-related/redis/service1/database"
	"github.com/go-related/redis/settings"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	appSettings settings.Service2
}

func InitService(appSettings settings.Service2) (*Service, error) {
	router := gin.Default()
	booksDb, err := database.NewBooks()
	if err != nil {
		log.WithError(err).Errorf("Setting up db.")
		return nil, err
	}
	redisService := redis2.New()

	handler := &Handler{
		booksDb,
		redisService,
	}
	handler.ListenForNewBook()

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
