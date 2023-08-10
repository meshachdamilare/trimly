package storage

import (
	"context"
	"github/meshachdamilare/trimly/model"
)

type RepoInterface interface {
	UrlRepoInterface
	UserRepoInterface
}

type UrlRepoInterface interface {
	Store(ctx context.Context, url *model.URL) error
	GetUrl(ctx context.Context, code string) (*model.URL, error)
	GetUrls(ctx context.Context, userId string) ([]model.URL, error)
	UpdateUrlVisitCount(ctx context.Context, url *model.URL) error
}

type UserRepoInterface interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByIdOrEmail(ctx context.Context, idOrEmail string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) error
	GetUserUrls(ctx context.Context, id string) ([]model.URL, error)
}
