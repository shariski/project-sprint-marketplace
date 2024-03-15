package service

import (
	"context"
	"project-sprint-marketplace/entity"
	"project-sprint-marketplace/model"
)

type PaymentService interface {
	Create(ctx context.Context, payment model.PaymentModel, userId int) entity.Payment
}
