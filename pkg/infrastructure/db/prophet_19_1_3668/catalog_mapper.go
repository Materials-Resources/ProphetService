package prophet_19_1_3668

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668/model"
)

func ToDBProduct(dbProduct *entities.Product) *model.InvMast {
	var im = &model.InvMast{
		ItemDesc:     dbProduct.Name,
		ExtendedDesc: sql.NullString{Valid: true, String: dbProduct.Description},
	}
	return im
}
func FromDBListProduct(dbProducts []model.InvLoc) ([]*entities.ValidatedProduct, error) {
	var validatedProducts []*entities.ValidatedProduct
	for _, product := range dbProducts {
		validatedProduct, err := entities.NewValidatedProduct(&entities.Product{
			ID:          strconv.Itoa(int(product.InvMastUid)),
			SN:          product.InvMast.ItemId,
			Name:        product.InvMast.ItemDesc,
			Description: product.InvMast.ExtendedDesc.String,
		},
		)
		if err != nil {
			fmt.Println(err)
		}
		validatedProducts = append(validatedProducts, validatedProduct)
	}
	return validatedProducts, nil
}
func FromDBProduct(dbProduct *model.InvMast) (*entities.ValidatedProduct, error) {
	var p = &entities.Product{Name: dbProduct.ItemDesc, Description: dbProduct.ExtendedDesc.String}
	p.ID = strconv.Itoa(int(dbProduct.InvMastUid))
	return entities.NewValidatedProduct(p)
}

func FromDBListGroup(dbProductGroups []model.ProductGroup) ([]*entities.ValidatedProductGroup, error) {
	var productGroups []*entities.ValidatedProductGroup

	for _, dbProductGroup := range dbProductGroups {
		validatedProductGroup, err := entities.NewValidatedProductGroup(&entities.ProductGroup{
			ID:   strconv.Itoa(int(dbProductGroup.ProductGroupUid)),
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

func FromDBProductGroup(dbProductGroup *model.ProductGroup) (*entities.ValidatedProductGroup, error) {
	var pg = &entities.ProductGroup{Name: dbProductGroup.ProductGroupDesc, SN: dbProductGroup.ProductGroupId}
	pg.ID = strconv.Itoa(int(dbProductGroup.ProductGroupUid))
	return entities.NewValidatedProductGroup(pg)
}
