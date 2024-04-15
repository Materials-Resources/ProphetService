package service

import (
	"context"
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_21_1_4559"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

func NewCatalogService(models prophet_21_1_4559.Models) Catalog {
	return Catalog{models}
}

type Catalog struct {
	models prophet_21_1_4559.Models
}

func (c Catalog) GetProduct(ctx context.Context, uid int32) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (c Catalog) CreateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error {

	dbProductGroup := &prophet_21_1_4559.ProductGroup{
		ProductGroupDesc: productGroup.Name, ProductGroupId: productGroup.SN,
	}

	err := c.models.ProductGroup.Insert(ctx, dbProductGroup)
	if err != nil {
		return err
	}

	productGroup.UID = dbProductGroup.ProductGroupUid

	return nil
}

func (c Catalog) GetProductGroup(ctx context.Context, sn string) (domain.ProductGroup, error) {
	dbProductGroup, err := c.models.ProductGroup.GetById(ctx, sn)
	if err != nil {
		return domain.ProductGroup{}, err
	}
	return domain.ProductGroup{
		UID: dbProductGroup.ProductGroupUid, SN: dbProductGroup.ProductGroupId, Name: dbProductGroup.ProductGroupDesc,
	}, err
}

func (c Catalog) UpdateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error {
	// select existing product group
	dbProductGroup, err := c.models.ProductGroup.GetById(ctx, productGroup.SN)
	if err != nil {
		return err
	}
	// update product group
	dbProductGroup.ProductGroupDesc = productGroup.Name

	// save updated product group
	err = c.models.ProductGroup.Update(ctx, dbProductGroup)
	if err != nil {
		return err
	}
	return nil
}

func (c Catalog) ListProductGroup(ctx context.Context) ([]*domain.ProductGroup, error) {
	dbProductGroups, err := c.models.ProductGroup.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	productGroups := make([]*domain.ProductGroup, len(dbProductGroups))
	for i, dbProductGroup := range dbProductGroups {
		productGroups[i] = &domain.ProductGroup{
			SN:   dbProductGroup.ProductGroupId,
			Name: dbProductGroup.ProductGroupDesc,
		}
	}
	return productGroups, nil
}
