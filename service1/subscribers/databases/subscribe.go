package databases

import (
	"context"
	"github.com/go-related/redis/service1/subscribers/model"
)

func (s *subscribeDB) Subscribe(ctx context.Context, data model.Subscribe) (model.Subscribe, error) {
	//TODO implement me
	panic("implement me")
}

func (s *subscribeDB) DeleteAllSubscribes(ctx context.Context, Id uint) error {
	//TODO implement me
	panic("implement me")
}

func (s *subscribeDB) GetAllSubscribes(ctx context.Context) ([]model.Subscribe, error) {
	//TODO implement me
	panic("implement me")
}

func (s *subscribeDB) GetSubscribeById(ctx context.Context, Id uint) (model.Subscribe, error) {
	//TODO implement me
	panic("implement me")
}
