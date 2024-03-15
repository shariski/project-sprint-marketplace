package repository

import (
	"context"
	"database/sql"
	"project-sprint-marketplace/entity"
)

type PaymentRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, payment entity.Payment) entity.Payment
}
