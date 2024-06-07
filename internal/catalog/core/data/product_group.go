package data

import (
	"context"
	"database/sql"
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

type productGroup struct {
	data.ProductGroup `bun:",extend"`
}

func (m *productGroup) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

func (m *productGroup) toProductGroupDomain(d *domain.ProductGroup) {
	d.Sn = m.ProductGroupId
	d.Name = &m.ProductGroupDesc

}

type productGroupDefaultValues struct {
	CompanyId        string
	AssetAccountNo   string
	DeleteFlag       string
	RevenueAccountNo string
	CosAccountNo     string
	LastMaintainedBy string
}

type ProductGroupModel struct {
	bun           bun.IDB
	defaultValues *productGroupDefaultValues
}

func NewProductGroupModel(bun bun.IDB) ProductGroupModel {
	return ProductGroupModel{
		bun: bun, defaultValues: &productGroupDefaultValues{
			CompanyId:        "MRS",
			AssetAccountNo:   "121001",
			DeleteFlag:       "N",
			RevenueAccountNo: "401001",
			CosAccountNo:     "501001",
			LastMaintainedBy: "admin",
		},
	}
}

func (m *ProductGroupModel) Get(ctx context.Context, sn string) (*domain.ProductGroup, error) {
	productGroupRec, err := m.get(ctx, sn)
	if err != nil {
		return nil, err
	}
	productGroupDomain := &domain.ProductGroup{}
	productGroupRec.toProductGroupDomain(productGroupDomain)
	return productGroupDomain, nil
}

type ProductGroupListOptions struct {
}

func (m *ProductGroupModel) List(ctx context.Context, opts *ProductGroupListOptions) ([]*domain.ProductGroup, error) {
	var pgs []productGroup
	q := m.bun.NewSelect().Model(&pgs).Where("company_id = ?", m.defaultValues.CompanyId).Order("product_group_id ASC")

	err := q.Scan(ctx)
	if err != nil {
		return nil, err
	}

	productGroups := make([]*domain.ProductGroup, len(pgs))
	for i, pg := range pgs {
		productGroups[i] = &domain.ProductGroup{}
		pg.toProductGroupDomain(productGroups[i])
	}
	return productGroups, nil
}

func (m *ProductGroupModel) Create(ctx context.Context, pg *domain.ProductGroup) error {
	pgData := productGroup{
		ProductGroup: data.ProductGroup{
			ProductGroupId:   pg.Sn,
			ProductGroupDesc: *pg.Name,
			CompanyId:        m.defaultValues.CompanyId,
			AssetAccountNo:   m.defaultValues.AssetAccountNo,
			DeleteFlag:       m.defaultValues.DeleteFlag,
			RevenueAccountNo: sql.NullString{String: m.defaultValues.RevenueAccountNo, Valid: true},
			CosAccountNo:     sql.NullString{String: m.defaultValues.CosAccountNo, Valid: true},
			LastMaintainedBy: m.defaultValues.LastMaintainedBy,
		},
	}
	_, err := m.bun.NewInsert().Model(&pgData).Exec(ctx)
	return err
}

func (m *ProductGroupModel) Update(ctx context.Context, pg *domain.ProductGroup) error {

	productGroupRec, err := m.get(ctx, pg.Sn)

	if err != nil {
		return err
	}

	if pg.Name != nil {
		productGroupRec.ProductGroupDesc = *pg.Name
	}
	_, err = m.bun.NewUpdate().Model(productGroupRec).WherePK().Exec(ctx)
	return err
}

func (m *ProductGroupModel) get(ctx context.Context, productGroupId string) (*productGroup, error) {
	var pg productGroup
	err := m.bun.NewSelect().Model(&pg).Where("company_id = ?", m.defaultValues.CompanyId).Where("product_group_id = ?", productGroupId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &pg, nil

}
