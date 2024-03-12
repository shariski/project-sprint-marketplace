package impl

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/common"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/repository"
)

type tagRepositoryImpl struct {
	*sql.DB
}

func NewTagRepositoryImpl(DB *sql.DB) repository.TagRepository {
	return &tagRepositoryImpl{DB: DB}
}

func (tagRepository *tagRepositoryImpl) Insert(ctx context.Context, tag entity.Tag) error {
	sql := `
		INSERT INTO tags (product_id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4);
	`

	err := tagRepository.DB.QueryRow(sql,
		&tag.ProductId, &tag.Name, common.GetDateNowUTCFormat(), common.GetDateNowUTCFormat()).Err()

	return err
}

func (tagRepository *tagRepositoryImpl) DeleteByProductId(ctx context.Context, productId int) error {
	sql := `
		DELETE FROM tags
		WHERE product_id = $1;
	`

	err := tagRepository.DB.QueryRow(sql, productId).Err()

	return err
}
