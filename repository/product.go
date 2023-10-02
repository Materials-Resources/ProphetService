package repository

import (
	"context"
	"strconv"

	"github.com/materials-resources/s_prophet/core/domain"
	"github.com/materials-resources/s_prophet/core/port/repository"
	"github.com/materials-resources/s_prophet/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
)

type productRepository struct {
	bun *bun.DB
	db  repository.Database
}

func (p productRepository) Read(ctx context.Context, id string) (*domain.Product, error) {
	invMast := new(model.InvMast)

	pk, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	err = p.bun.NewSelect().Model(invMast).Where(
		"inv_mast_uid = ?",
		pk,
	).Scan(ctx)
	if err != nil {
		return nil, err
	}
	product := invMastToDomain(invMast)

	return &product, nil

}

func (p productRepository) Create(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Update(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Delete(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db repository.Database) repository.ProductRepository {
	bundb := bun.NewDB(
		db.GetDB(),
		mssqldialect.New(),
	)
	return &productRepository{
		db:  db,
		bun: bundb,
	}
}

func invMastToDomain(invMast *model.InvMast) domain.Product {
	product := domain.Product{
		Name:        invMast.ItemDesc,
		Description: invMast.ExtendedDesc.String,
		SN:          invMast.ItemId,
		ID:          strconv.Itoa(int(invMast.InvMastUid)),
		Price:       0,
	}
	return product
}
