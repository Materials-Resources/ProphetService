package service

import (
	"context"
	"database/sql"
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_21_1_4559"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

func NewCatalogService(models prophet_21_1_4559.Models) Catalog {
	return Catalog{models}
}

type Catalog struct {
	models prophet_21_1_4559.Models
}

func (c Catalog) UpdateProduct(ctx context.Context, product *domain.Product, locations []float64) error {
	invMast, err := c.models.InvMast.Get(ctx, product.UID)
	if err != nil {
		return err
	}

	if product.Name != "" {
		invMast.ItemDesc = product.Name
	}
	if product.Description != "" {
		invMast.ExtendedDesc = sql.NullString{String: product.Description, Valid: true}
	}

	err = c.models.InvMast.Update(ctx, invMast)
	if err != nil {
		return err
	}
	invLocs, err := c.models.InvLoc.GetByInvMastUid(ctx, locations, product.UID)
	if err != nil {
		return err
	}
	for _, invLoc := range invLocs {
		if product.ProductGroupSn != "" {
			invLoc.ProductGroupId = sql.NullString{String: product.ProductGroupSn, Valid: true}
		}
		err := c.models.InvLoc.Update(ctx, &invLoc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c Catalog) GetProductBySupplierPartNumber(
	ctx context.Context, partNumber string,
	supplier float64) (
	domain.Product,
	error) {
	invMast, err := c.models.InvMast.GetBySupplierPartNumber(ctx, supplier, partNumber)
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		UID: invMast.InvMastUid, SN: invMast.ItemId, Name: invMast.ItemDesc,
	}, nil
}

func (c Catalog) ListProduct(ctx context.Context, filter domain.Filter) ([]*domain.Product, error) {
	dbFilter := &prophet_21_1_4559.Filters{Cursor: filter.Cursor, Limit: filter.Limit}

	invLocs, _, err := c.models.InvLoc.GetAll(ctx, 1001, prophet_21_1_4559.DeleteFlagNo, *dbFilter)

	products := make([]*domain.Product, len(invLocs))
	for i, invLoc := range invLocs {
		products[i] = &domain.Product{
			UID: invLoc.InvMastUid, SN: invLoc.InvMast.ItemId, Name: invLoc.InvMast.ItemDesc,
			Description: invLoc.InvMast.ExtendedDesc.String,
		}
	}

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (c Catalog) GetProduct(ctx context.Context, uid int32) (domain.Product, error) {
	// TODO implement me
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

func (c Catalog) GetProductGroup(ctx context.Context, sn string) (domain.ProductGroup, []*domain.Product, error) {
	dbProductGroup, err := c.models.ProductGroup.GetById(ctx, sn)
	if err != nil {
		return domain.ProductGroup{}, nil, err
	}

	dbInvLocs, err := c.models.InvLoc.GetByProductGroupId(ctx, sn)

	products := make([]*domain.Product, len(dbInvLocs))

	for i, dbInvLoc := range dbInvLocs {
		products[i] = &domain.Product{
			UID: dbInvLoc.InvMastUid, SN: dbInvLoc.InvMast.ItemId, Name: dbInvLoc.InvMast.ItemDesc,
			Description: dbInvLoc.InvMast.ExtendedDesc.String,
		}

	}

	return domain.ProductGroup{
		UID: dbProductGroup.ProductGroupUid, SN: dbProductGroup.ProductGroupId, Name: dbProductGroup.ProductGroupDesc,
	}, products, nil
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
