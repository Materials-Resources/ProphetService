package repository

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
)

type productGroupDefaultValues struct {
	CompanyId        string
	AssetAccountNo   string
	DeleteFlag       string
	RevenueAccountNo string
	CosAccountNo     string
	LastMaintainedBy string
}

// NewProductGroupRepository creates a new instance of ProductGroupRepository
func NewProductGroupRepository(db bun.IDB) *ProductGroupRepository {
	return &ProductGroupRepository{
		db: db,
		defaultValues: &productGroupDefaultValues{
			CompanyId:        "MRS",
			AssetAccountNo:   "121001",
			DeleteFlag:       "N",
			RevenueAccountNo: "401001",
			CosAccountNo:     "501001",
			LastMaintainedBy: "admin",
		},
	}
}

type ProductGroupRepository struct {
	db            bun.IDB
	defaultValues *productGroupDefaultValues
}

// GetProductGroup returns a productGroup with its details
func (r *ProductGroupRepository) GetProductGroup(ctx context.Context, id string) (*domain.ProductGroup, error) {
	var productGroup productGroup
	err := r.db.NewSelect().Model(&productGroup).Where("product_group_id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return mapDbProductGroupToDomainProductGroup(&productGroup), nil
}

// ListProductGroups returns a list of ProductGroups
func (r *ProductGroupRepository) ListProductGroups(ctx context.Context) ([]*domain.ProductGroup, error) {
	var productGroups []*productGroup
	err := r.db.NewSelect().Model(&productGroups).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var domainProductGroups []*domain.ProductGroup
	for _, productGroup := range productGroups {
		domainProductGroups = append(domainProductGroups, mapDbProductGroupToDomainProductGroup(productGroup))
	}
	return domainProductGroups, nil
}

func (r *ProductGroupRepository) CreateProductGroup(ctx context.Context, input *domain.ProductGroup) error {
	pgData := productGroup{
		ProductGroup: prophet.ProductGroup{
			ProductGroupId:   input.Sn,
			ProductGroupDesc: *input.Name,
			CompanyId:        r.defaultValues.CompanyId,
			AssetAccountNo:   r.defaultValues.AssetAccountNo,
			DeleteFlag:       r.defaultValues.DeleteFlag,
			RevenueAccountNo: &r.defaultValues.RevenueAccountNo,
			CosAccountNo:     &r.defaultValues.CosAccountNo,
			LastMaintainedBy: r.defaultValues.LastMaintainedBy,
		},
	}
	_, err := r.db.NewInsert().Model(&pgData).Exec(ctx)
	return err
}

func (r *ProductGroupRepository) UpdateProductGroup(ctx context.Context, group *domain.ProductGroup) error {
	var productGroupRec productGroup
	err := r.db.NewSelect().Model(&productGroupRec).Where("product_group_id = ?", group.Sn).Scan(ctx)

	if group.Name != nil {
		productGroupRec.ProductGroupDesc = *group.Name
	}
	_, err = r.db.NewUpdate().Model(&productGroupRec).ExcludeColumn("product_group_uid").WherePK().Exec(ctx)
	return err
}

func mapDbProductGroupToDomainProductGroup(productGroup *productGroup) *domain.ProductGroup {
	return &domain.ProductGroup{
		Sn:   productGroup.ProductGroupId,
		Name: &productGroup.ProductGroupDesc,
	}
}
