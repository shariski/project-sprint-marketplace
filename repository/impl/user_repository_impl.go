package impl

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/repository"
)

type userRepositoryImpl struct {
	*sql.DB
}

func NewUserRepositoryImpl(DB *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

func (userRepository *userRepositoryImpl) Insert(ctx context.Context, user entity.User) entity.User {
	sql := `
		INSERT INTO users (username, name, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING username, name;
		`
	var insertedUser entity.User
	err := userRepository.DB.QueryRow(sql, &user.Username, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt).Scan(&insertedUser.Username, &insertedUser.Name)
	if err != nil {
		panic(err)
	}
	return insertedUser
}
