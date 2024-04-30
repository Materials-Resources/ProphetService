package service

import (
	"context"
	"github.com/materials-resources/s_prophet/internal/customer/domain"
)

type Service interface {
	GetContactById(ctx context.Context, id string) (domain.Contact, error)
}
