package repository

import (
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_21_1_4559"
	"github.com/materials-resources/s_prophet/internal/inventory/domain"
)

type invLoc prophet_21_1_4559.InvLoc

func (i *invLoc) WriteToProductStock() *domain.ProductStock {
	return &domain.ProductStock{
		QuantityAvailable: i.QtyOnHand.Float64,
		ProductUID:        i.InvMastUid,
	}
}
