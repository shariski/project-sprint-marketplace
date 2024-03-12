package impl

import (
	"context"
	"database/sql"
	"errors"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/repository"
	"strings"
)

type userRepositoryImpl struct {
	*sql.DB
}

func NewUserRepositoryImpl(DB *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

func (userRepository *userRepositoryImpl) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	sql := `
		INSERT INTO users (username, name, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING username, name;
	`
	var insertedUser entity.User
	if err := userRepository.DB.QueryRow(sql,
		&user.Username, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt).Scan(&insertedUser.Username, &insertedUser.Name); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return entity.User{}, errors.New("duplicate key")
		}
	}
	return insertedUser, nil
}
