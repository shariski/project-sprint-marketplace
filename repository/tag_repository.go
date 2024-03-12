package repository

import (
	"context"
	"project-sprint-marketplace/entity"
)

type TagRepository interface {
	Insert(ctx context.Context, tag entity.Tag) error
	DeleteByProductId(ctx context.Context, id int) error
}