package repository

import (
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	"github.com/materials-resources/s_prophet/pkg/internal/infra/repository/model"
)

type productRepository struct {
	db repository.Database
}

func (p productRepository) Read(id string) (*domain.Product, error) {
	var invMast model.InvMast

	pk, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	err = p.db.GetDB().Model(&model.InvMast{}).First(
		&invMast,
		pk,
	).Error
	if err != nil {
		return nil, err
	}
	product := invMastToDomain(&invMast)

	return &product, nil

}

func (p productRepository) Create() {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Update() {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Delete() {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db repository.Database) repository.ProductRepository {
	return &productRepository{db: db}
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
