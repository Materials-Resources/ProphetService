package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

func ProductSupplierFromDomain(d *domain.ProductSupplier, m *prophet_19_1_3668.InventorySupplier) error {
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
			&domain.Product{Name: m.InvMast.ItemDesc, Description: m.InvMast.ExtendedDesc.String,
				ID: strconv.Itoa(int(m.InvMastUid)),
				SN: m.InvMast.ItemId, ProductGroupName: m.ProductGroup.ProductGroupDesc,
			},
		)
	}
	return ds
}
func ToDBProduct(dbProduct *domain.Product) *prophet_19_1_3668.InvMast {
	var im = &prophet_19_1_3668.InvMast{
		ItemDesc:     dbProduct.Name,
		ExtendedDesc: sql.NullString{Valid: true, String: dbProduct.Description},
	}
	return im
}

//	func FromDBListProduct(dbProducts []prophet_19_1_3668.InvLoc) ([]*domain.ValidatedProduct, error) {
//		var validatedProducts []*domain.ValidatedProduct
//		for _, product := range dbProducts {
//			validatedProduct, err := domain.NewValidatedProduct(&domain.Product{
//				ID:          strconv.Itoa(int(product.InvMastUid)),
//				SN:          product.InvMast.ItemId,
//				Name:        product.InvMast.ItemDesc,
//				Description: product.InvMast.ExtendedDesc.String,
//			},
//			)
//			if err != nil {
//				fmt.Println(err)
//			}
//			validatedProducts = append(validatedProducts, validatedProduct)
//		}
//		return validatedProducts, nil
//	}
func FromDBProduct(dbProduct *prophet_19_1_3668.InvMast) (*domain.ValidatedProduct, error) {
	var p = &domain.Product{Name: dbProduct.ItemDesc, Description: dbProduct.ExtendedDesc.String}
	p.ID = strconv.Itoa(int(dbProduct.InvMastUid))
	return domain.NewValidatedProduct(p)
}

func FromDBListGroup(dbProductGroups []prophet_19_1_3668.ProductGroup) ([]*domain.ValidatedProductGroup, error) {
	var productGroups []*domain.ValidatedProductGroup

	for _, dbProductGroup := range dbProductGroups {
		validatedProductGroup, err := domain.NewValidatedProductGroup(&domain.ProductGroup{
			Name: dbProductGroup.ProductGroupDesc,
			SN:   dbProductGroup.ProductGroupId,
		},
		)
		if err != nil {
			fmt.Println(err)
		}
		productGroups = append(productGroups, validatedProductGroup)
	}
	return productGroups, nil
}

func FromDBProductGroup(dbProductGroup *prophet_19_1_3668.ProductGroup) (*domain.ValidatedProductGroup, error) {
	var pg = &domain.ProductGroup{Name: dbProductGroup.ProductGroupDesc, SN: dbProductGroup.ProductGroupId}
	return domain.NewValidatedProductGroup(pg)
}
