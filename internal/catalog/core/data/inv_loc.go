package data

import (
	"context"
	"errors"
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/uptrace/bun"
	"sort"
	"strconv"
	"sync"
	"time"
)

type invLoc struct {
	data.InvLoc  `bun:",extend"`
	InvMast      *invMast      `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
	ProductGroup *productGroup `bun:"rel:has-one,join:product_group_id=product_group_id"`
}

var _ bun.BeforeAppendModelHook = (*invLoc)(nil)

func (m *invLoc) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		m.LastMaintainedBy = "admin"
		m.DateLastModified = time.Now()
	}
	return nil
}

type invLocDefaultValues struct {
	CompanyId  string
	LocationId float64
}

func (m *invLoc) toProductDomain(d *domain.Product) {
	if m.InvMast != nil {
		d.Uid = strconv.Itoa(int(m.InvMast.InvMastUid))
		d.Sn = &m.InvMast.ItemId
		d.Name = &m.InvMast.ItemDesc
		d.Description = &m.InvMast.ItemDesc
		d.ProductGroupSn = &m.ProductGroupId.String
	}

	if m.ProductGroup != nil {
		d.ProductGroupName = m.ProductGroup.ProductGroupDesc
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

func (m *InvLocModel) List(ctx context.Context, opts *InvLocListOptions, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error) {
	var invLocRecs []*invLoc
	var totalCount int

	q := m.bun.NewSelect().Model(&invLocRecs).
		Relation("InvMast").Relation("ProductGroup").Limit(int(pagination.Size))

	// TODO create a persistent var of q for the count query
	var wg sync.WaitGroup

	q, err := m.listFilterQuery(q, opts)

	go func() {
		wg.Add(1)
		defer wg.Done()
		c, err := q.Count(ctx)
		if err != nil {
			// TODO handle error
			return
		}
		totalCount = c
	}()

	q, err = m.listPaginationQuery(q, pagination)

	err = q.Scan(ctx)
	if err != nil {
		return nil, m.calculatePaginationMetadata(ctx, int64(totalCount), invLocRecs), err
	}

	sort.Slice(invLocRecs, func(i, j int) bool {
		return invLocRecs[i].InvMastUid < invLocRecs[j].InvMastUid
	})

	productDomains := make([]*domain.Product, len(invLocRecs))

	for i, rec := range invLocRecs {
		productDomains[i] = &domain.Product{}
		rec.toProductDomain(productDomains[i])
	}

	wg.Wait()

	return productDomains, m.calculatePaginationMetadata(ctx, int64(totalCount), invLocRecs), nil
}

func (m *InvLocModel) listFilterQuery(q *bun.SelectQuery, opts *InvLocListOptions) (*bun.SelectQuery, error) {
	if opts.ProductGroupIds != nil {
		q = q.Where("product_group_id IN (?)", bun.In(*opts.ProductGroupIds))
	}
	return q, nil
}

func (m *InvLocModel) listPaginationQuery(q *bun.SelectQuery, pagination domain.Pagination) (*bun.SelectQuery, error) {
	switch pagination.Direction {
	case domain.CursorDirectionPrevious:
		q = q.Where("inv_loc.inv_mast_uid < ?", pagination.Cursor).
			OrderExpr("inv_loc.inv_mast_uid DESC")
	case domain.CursorDirectionNext:
		q = q.Where("inv_loc.inv_mast_uid > ?", pagination.Cursor).
			OrderExpr("inv_loc.inv_mast_uid ASC")
	default:
		return nil, errors.New("invalid cursor direction")
	}
	return q, nil
}

func (m *InvLocModel) Update(ctx context.Context, product *domain.Product) error {
	uidInt, err := strconv.Atoi(product.Uid)
	if err != nil {
		// TODO return not found error
		return err
	}
	rec, err := m.get(ctx, int32(uidInt))

	if err != nil {
		return err
	}

	q := m.bun.NewUpdate().Model(rec).WherePK()

	if product.ProductGroupSn != nil {
		rec.ProductGroupId.String = *product.ProductGroupSn
	}
	_, err = q.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *InvLocModel) get(ctx context.Context, invMastUid int32) (*invLoc, error) {
	var invLocRec invLoc
	err := m.bun.NewSelect().Model(&invLocRec).
		Where("inv_mast_uid = ?", invMastUid).
		Where("company_id = ?", m.defaultValues.CompanyId).
		Where("location_id = ?", m.defaultValues.LocationId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &invLocRec, nil
}

func (m *InvLocModel) calculatePaginationMetadata(ctx context.Context, total int64, recs []*invLoc) *domain.PaginationMetadata {
	paginationMetadata := &domain.PaginationMetadata{
		NextCursor:     0,
		PreviousCursor: 0,
		Total:          total,
	}

	if len(recs) > 0 {
		paginationMetadata.NextCursor = recs[len(recs)-1].InvMastUid
		paginationMetadata.PreviousCursor = recs[0].InvMastUid
	}

	return paginationMetadata

}
