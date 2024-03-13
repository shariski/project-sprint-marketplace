package impl

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/repository"
)

type tagRepositoryImpl struct {}

func NewTagRepositoryImpl() repository.TagRepository {
	return &tagRepositoryImpl{}
}

func (tagRepository *tagRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, tag entity.Tag) entity.Tag {
	SQL := `
		INSERT INTO tags (product_id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4);
	`

	timeNow := common.GetDateNowUTCFormat()

	_, err := tx.ExecContext(ctx, SQL,
		&tag.ProductId, &tag.Name, timeNow, timeNow)

	exception.PanicLogging(err)

	tag.CreatedAt = timeNow
	tag.UpdatedAt = timeNow

	return tag
}

func (tagRepository *tagRepositoryImpl) DeleteByProductId(ctx context.Context, tx *sql.Tx, productId int) {
	SQL := `
		DELETE FROM tags
		WHERE product_id = $1;
	`

	_, err := tx.ExecContext(ctx, SQL, productId)

	exception.PanicLogging(err)
}
