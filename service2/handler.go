package service2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-related/redis/redis"
	"github.com/go-related/redis/service1/books/datebase"
	"github.com/go-related/redis/service1/books/model"
	"github.com/go-related/redis/service1/subscribers/databases"
	"github.com/go-related/redis/settings"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	BookDb       datebase.BooksDB
	SubscriberDB databases.SubscribeDB
	Redis        *redis.RedisService
}

func (h *Handler) ListenForNewBook() {
	subscriber := h.Redis.SubscribeToChannel(context.Background(), settings.ApplicationConfiguration.Service2.NewBookChannelName)

	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			logrus.WithError(err).Error("error receiving message")
		}
		book := model.Book{}
		if err := json.Unmarshal([]byte(msg.Payload), &book); err != nil {
			logrus.WithError(err).Error("error parsing book from channel")
			continue
		}
		go h.sendBooksEmail(book)
	}

}

func (h *Handler) sendBooksEmail(book model.Book) {
	data, err := h.BookDb.GetBookById(context.Background(), book.ID)
	if err != nil {
		logrus.WithError(err).Error("error loading book")
		return
	}
	if !data.IsDeleted { //check for sanity
		subscribers, err := h.SubscriberDB.GetAuthorsSubscribers(context.Background(), data.Authors)
		if err != nil {
			logrus.WithError(err).Error("error loading subscribers")
			return
		}
		subject := "New Book !!!"
		for _, subscriber := range subscribers {
			body := fmt.Sprintf("Hi %s,\n New book '%s' is out.\n Enjoy!!!!", subscriber.FullName, book.Title)
			go h.sendEmail(body, subject, subscriber.Email)
		}
	}
}

func (h *Handler) sendEmail(body string, subject string, email string) {
	logrus.WithField("subject", subject).WithField("body", body).WithField("email", email).Info("sending email to person")
	// here will be the actual implementation oto send the email but for this we won't implement this
}
