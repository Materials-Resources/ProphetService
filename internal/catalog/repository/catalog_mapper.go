package repository

import (
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

func FromDBListProduct(ms []prophet_19_1_3668.InvLoc) []*domain.Product {
	var ds []*domain.Product
	for _, m := range ms {
		ds = append(
			ds,
			&domain.Product{
				Name: m.InvMast.ItemDesc, Description: m.InvMast.ExtendedDesc.String,
				UID: m.InvMastUid,
			},
		)
	}
	return ds
}
