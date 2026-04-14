package user

import (
	"context"

	"github.com/aarondl/authboss/v3"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {

	return &UserService{repo: repo}
}

func (service *UserService) Load(ctx context.Context, key string) (authboss.User, error) {

	return service.repo.Load(ctx, key)
}

func (service *UserService) Save(ctx context.Context, user authboss.User) error {

	return service.repo.Save(ctx, user)

}

func (service *UserService) New(ctx context.Context) authboss.User {

	return service.repo.New(ctx)

}

func (service *UserService) Create(ctx context.Context, user authboss.User) error {

	return service.repo.Create(ctx, user)

}
