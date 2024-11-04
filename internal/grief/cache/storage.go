package cache

import (
	"context"

	"github.com/Wexlersolk/GriefBlades/internal/grief"
	"github.com/go-redis/redis"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*grief.User, error)
		Set(context.Context, grief.User) error
		Delete(context.Context, int64)
	}
}

func NewRedisStorage(rbd *redis.Client) Storage {
	return Storage{
		Users: &UserStore{rdb: rbd},
	}
}
