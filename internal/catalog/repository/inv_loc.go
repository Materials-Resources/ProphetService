package repository

import (
	"context"
	"errors"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/materials-resources/s-prophet/pkg/helpers"
	"github.com/uptrace/bun"
	"sort"
	"strconv"
	"sync"
	"time"
)

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
		d.Id = strconv.Itoa(int(m.InvMast.InvMastUid))
		d.Sn = m.InvMast.ItemId
		d.Name = m.InvMast.ItemDesc
		d.Description = helpers.GetValueOrDefault(m.InvMast.ExtendedDesc, "")
		d.ProductGroupSn = helpers.GetValueOrDefault(m.ProductGroupId, "")
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

	rec, err := m.get(ctx, int32(uidInt))
	if err != nil {
		return nil, err
	}

	productDomain := &domain.Product{}
	rec.toProductDomain(productDomain)

	return productDomain, nil
}

type InvLocListOptions struct {
	ProductGroupIds *[]string
}

func (m *InvLocModel) List(ctx context.Context, opts *InvLocListOptions, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error) {
	var invLocRecs []*invLoc
	var totalCount int

	q := m.bun.NewSelect().Model(&invLocRecs).
		Relation("InvMast").Relation("productGroup").Limit(int(pagination.Size))

	var wg sync.WaitGroup

	q, err := m.listFilterQuery(q, opts)
	if err != nil {
		return nil, nil, err
	}

	countQ := *q

	wg.Add(1)
	go func() {
		defer wg.Done()
		c, err := countQ.Count(ctx)
		if err != nil {
			// TODO handle error
		}
		totalCount = c
	}()

	q, err = m.listPaginationQuery(q, pagination)
	if err != nil {
		return nil, nil, err

	}

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

func (m *InvLocModel) get(ctx context.Context, invMastUid int32) (*invLoc, error) {
	var invLocRec invLoc
	err := m.bun.NewSelect().
		Model(&invLocRec).
		Relation("InvMast").
		Where("inv_loc.inv_mast_uid = ?", invMastUid).
		Where("inv_loc.company_id = ?", m.defaultValues.CompanyId).
		Where("inv_loc.location_id = ?", m.defaultValues.LocationId).Scan(ctx)
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
