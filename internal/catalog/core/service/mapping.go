package service

import (
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"strconv"
)

func mapInvLocToProduct(invLoc *data.InvLoc, product *domain.Product) {
	product.Uid = strconv.Itoa(int(invLoc.InvMastUid))
	product.Sn = &invLoc.InvMast.ItemId
	product.Name = &invLoc.InvMast.ItemDesc
	product.Description = &invLoc.InvMast.ExtendedDesc.String
	product.ProductGroupName = invLoc.ProductGroupId.String
}
