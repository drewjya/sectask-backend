package entity

import (
	"context"
)

type Repository interface {
	GetEntity(ctx context.Context, entityID string) (err error)
}
