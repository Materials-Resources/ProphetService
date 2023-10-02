package repository

import (
	"context"

	"github.com/materials-resources/s_prophet/core/domain"
)

type ReceivingRepository interface {
	ReadReceipt(ctx context.Context, id string) (*domain.ReceivingReceipt, error)
}
