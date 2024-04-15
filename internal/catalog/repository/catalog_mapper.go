package repository

import (
	"database/sql"
	"strconv"

	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

func ProductSupplierFromDomain(d domain.ProductSupplier, m *inventorySupplier) error {
	productId, err := strconv.Atoi(d.ProductId)
	if err != nil {
		return err
	}
	supplierId, err := strconv.ParseFloat(d.SupplierId, 64)
	if err != nil {
		return err
	}
	m.InvMastUid = int32(productId)
	m.SupplierId = supplierId
	m.DivisionId = supplierId
	m.SupplierPartNo = sql.NullString{String: d.SupplierSn, Valid: true}
	m.ListPrice = d.ListPrice
	m.Cost = d.PurchasePrice

	return nil

}

func FromDBListProduct(ms []prophet_19_1_3668.InvLoc) []*domain.Product {
	var ds []*domain.Product
	for _, m := range ms {
		ds = append(ds,
			&domain.Product{
				Name: m.InvMast.ItemDesc, Description: m.InvMast.ExtendedDesc.String,
				ID: m.InvMastUid,
			},
		)
	}
	return ds
}
func ToDBProduct(dbProduct *domain.Product) *invMast {
	var im = &invMast{
		ItemDesc:     dbProduct.Name,
		ExtendedDesc: sql.NullString{Valid: true, String: dbProduct.Description},
	}
	return im
}
