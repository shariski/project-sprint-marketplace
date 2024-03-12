package repository

import (
	"context"
	"project-sprint-marketplace/entity"
)

type UserRepository interface {
	// Authentication(ctx context.Context, username string, password string) (entity.User, error)
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	// FindByUsername(ctx context.Context, username string) entity.User
}
