package redis

import "context"

type Repository interface {
	Get(key string, cb interface{}) error
	SetWithExp(ctx context.Context, key string, val interface{}, duration string) error
}
