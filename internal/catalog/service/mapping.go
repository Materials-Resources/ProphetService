package service

import (
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"strconv"
)

func mapInvLocToProduct(invLoc *data.InvLoc, product *domain.Product) {
	product.Uid = strconv.Itoa(int(invLoc.InvMastUid))
	product.Sn = invLoc.InvMast.ItemId
	product.Name = invLoc.InvMast.ItemDesc
	product.Description = invLoc.InvMast.ExtendedDesc.String
	product.ProductGroupName = invLoc.ProductGroupId.String
	product.ListPrice = invLoc.Price1.Float64
}

func CreateProductGroupFromRecord(productGroup *data.ProductGroup) *domain.ProductGroup {
	return &domain.ProductGroup{
		SN:   productGroup.ProductGroupId,
		Name: productGroup.ProductGroupDesc,
	}
}
func CreateProductFromRecord(invLoc *data.InvLoc) *domain.Product {
	product := &domain.Product{
		Uid:              strconv.Itoa(int(invLoc.InvMastUid)),
		Sn:               invLoc.InvMast.ItemId,
		Name:             invLoc.InvMast.ItemDesc,
		Description:      invLoc.InvMast.ExtendedDesc.String,
		ProductGroupId:   invLoc.ProductGroupId.String,
		ProductGroupName: invLoc.ProductGroup.ProductGroupDesc,
		ListPrice:        0,
		StockQuantity:    0,
	}
	return product
}
