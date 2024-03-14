package repository

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/entity"
)

type TagRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, tag entity.Tag) entity.Tag
	DeleteByProductId(ctx context.Context, tx *sql.Tx, id int)
}
