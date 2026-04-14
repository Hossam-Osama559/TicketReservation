package user

import (
	"context"

	"github.com/aarondl/authboss/v3"
)

type UserRepository interface {
	Load(ctx context.Context, key string) (authboss.User, error)

	Save(ctx context.Context, user authboss.User) error

	New(ctx context.Context) authboss.User

	Create(ctx context.Context, user authboss.User) error
}
