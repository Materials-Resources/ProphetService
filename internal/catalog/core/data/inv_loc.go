package data

import (
	"context"
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/uptrace/bun"
	"strconv"
)

type invLoc struct {
	data.InvLoc `bun:",extend"`
	InvMast     *invMast `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
}

type invLocDefaultValues struct {
	CompanyId  string
	LocationId float64
}

func (l *invLoc) toProductDomain(d *domain.Product) {
	if l.InvMast != nil {
		d.Uid = strconv.Itoa(int(l.InvMast.InvMastUid))
		d.Sn = l.InvMast.ItemId
		d.Name = l.InvMast.ItemDesc
		d.Description = l.InvMast.ItemDesc
		d.ProductGroupId = l.ProductGroupId.String
	}

}

type InvLocModel struct {
	bun           bun.IDB
	defaultValues *invLocDefaultValues
}

func NewInvLocModel(bun bun.IDB) InvLocModel {
	return InvLocModel{
		bun: bun, defaultValues: &invLocDefaultValues{
			CompanyId:  "MRS",
			LocationId: 1001,
		},
	}
}

func (m *InvLocModel) Get(ctx context.Context, uid string) (*domain.Product, error) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		// TODO return not found error
		return nil, err
	}
	var invLocRec invLoc
	err = m.bun.NewSelect().Model(&invLocRec).
		Relation("InvMast").
		Where("inv_loc.company_id = ? and inv_loc.location_id = ?", m.defaultValues.CompanyId, m.defaultValues.LocationId).
		Where("inv_loc.inv_mast_uid = ?", uidInt).Scan(ctx)
	if err != nil {
		return nil, err
	}

	productDomain := &domain.Product{}
	invLocRec.toProductDomain(productDomain)

	return productDomain, nil
}

type InvLocListOptions struct {
	ProductGroupIds *[]string
}

func (m *InvLocModel) List(ctx context.Context, opts *InvLocListOptions) ([]*domain.Product, error) {
	var invLocRecs []*invLoc

	q := m.bun.NewSelect().Model(&invLocRecs).
		Relation("InvMast")

	if opts.ProductGroupIds != nil {
		q.Where("product_group_id IN (?)", bun.In(*opts.ProductGroupIds))
	}

	err := q.Scan(ctx)
	if err != nil {
		return nil, err
	}

	productDomains := make([]*domain.Product, len(invLocRecs))

	for i, rec := range invLocRecs {
		productDomains[i] = &domain.Product{}
		rec.toProductDomain(productDomains[i])
	}

	return productDomains, nil
}
