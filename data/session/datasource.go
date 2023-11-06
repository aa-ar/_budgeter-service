package session

import (
	"context"

	"github.com/aa-ar/budgeter-service/internal/data/redis"
	rdb "github.com/redis/go-redis/v9"
)

type SessionDataSource struct {
	*rdb.Client
	ctx context.Context
}

func NewSessionDataSource(config *redis.Config) *SessionDataSource {
	ctx := context.Background()
	return &SessionDataSource{
		Client: config.Connect(),
		ctx:    ctx,
	}
}
