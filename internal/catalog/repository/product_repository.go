package repository

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/materials-resources/s-prophet/pkg/helpers"
	"github.com/uptrace/bun"
	"strconv"
)

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db bun.IDB, defaultValues DefaultProductValues) *ProductRepository {
	return &ProductRepository{db: db, defaultValues: &defaultValues}
}

type DefaultProductValues struct {
	CompanyId        string
	LocationId       string
	GlAccountNo      string
	RevenueAccountNo string
	CosAccountNo     string
}

type ProductRepository struct {
	db            bun.IDB
	defaultValues *DefaultProductValues
}

// GetProduct returns a Product with its details
func (r *ProductRepository) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	var record invLoc

	err := r.db.NewSelect().Model(&record).Relation("InvMast").Where("inv_loc.inv_mast_uid = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	product := domain.Product{
		Id:             strconv.Itoa(int(record.InvMastUid)),
		Sn:             record.InvMast.ItemId,
		Name:           record.InvMast.ItemDesc,
		Description:    helpers.GetOptionalValue(record.InvMast.ExtendedDesc, ""),
		ProductGroupSn: helpers.GetValueOrDefault(record.ProductGroupId, ""),
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, update *domain.ProductUpdate) error {

	invMastUid, err := strconv.Atoi(update.Uid)
	if err != nil {
		return err
	}
	invMastRecord, err := r.selectInvMastRecord(ctx, invMastUid)
	if err != nil {
		return err
	}

	if update.Name != nil {
		invMastRecord.ItemDesc = *update.Name
	}
	if update.Description != nil {
		invMastRecord.ExtendedDesc = update.Description
	}

	if update.ProductGroupSn != nil {
		for _, invLoc := range invMastRecord.InvLocs {
			invLoc.ProductGroupId = update.ProductGroupSn
		}
	}

	_, err = r.db.NewUpdate().Model(&invMastRecord).WherePK().Exec(ctx)

	if err != nil {
		return err
	}

	_, err = r.db.NewUpdate().Model(&invMastRecord.InvLocs).WherePK().Exec(ctx)

	return nil

}

func (r *ProductRepository) UpdatePrimarySupplier(ctx context.Context, update *domain.ProductUpdate) error {
	//err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
	//
	//	inventorySupplier := &inventorySupplier{}
	//	err := tx.NewSelect().Model(inventorySupplier).Scan(ctx)
	//})

	return nil

}

// ListProducts returns a list of Products
func (r *ProductRepository) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	var records []*InvMast
	err := r.db.NewSelect().Model(&records).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return nil, err
}

// DeleteProduct deletes a Product by its ID
func (r *ProductRepository) DeleteProduct(ctx context.Context, uid string) error {

	invMastUid, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}

	ok, err := r.safeToDeleteProduct(ctx, invMastUid)

	if err != nil || !ok {
		return err
	}

	var record InvMast

	err = r.db.NewSelect().Model(&record).
		Relation("AlternateCodes").
		Relation("assemblyHdr").
		Relation("assemblyHdr.AssemblyLines").
		Relation("AverageInventoryValues").
		Relation("InventorySuppliers").
		Relation("InventorySuppliers.InventorySupplierXLocs").
		Relation("InventorySuppliers.InventorySupplierXUoms").
		Relation("InvLocs").
		Relation("InvLocs.InvBins").
		Relation("InvLocs.InvTrans").
		Relation("InvLocs.InvLocStockStatuses").
		Relation("InvLocMsps").
		Relation("InvXrefs").
		Relation("ItemIdChangeHistories").
		Relation("ItemCategoryXInvMasts").
		Relation("ItemUoms").
		Relation("PricePages").
		Relation("PricePages.PricePageXBooks").
		Where("inv_mast.inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return err
	}

	if record.InvLocs != nil {
		for _, invLoc := range record.InvLocs {
			if invLoc.InvBins != nil {
				for _, invBin := range invLoc.InvBins {
					_, err := r.db.NewDelete().Model(invBin).WherePK().Exec(ctx)
					if err != nil {
						return err
					}
				}
			}
			if invLoc.InvLocStockStatuses != nil {
				for _, invLocStockStatus := range invLoc.InvLocStockStatuses {
					_, err := r.db.NewDelete().Model(invLocStockStatus).WherePK().Exec(ctx)
					if err != nil {
						return err
					}
				}
			}
			_, err := r.db.NewDelete().Model(invLoc).WherePK().Exec(ctx)
			if err != nil {
				return err
			}
		}

	}

	if record.AlternateCodes != nil {
		for _, altCode := range record.AlternateCodes {
			_, err := r.db.NewDelete().Model(altCode).WherePK().Exec(ctx)
			if err != nil {
				return err
			}
		}

	}

	// delete assembly_hdr records
	if record.AssemblyHdr != nil {
		if record.AssemblyHdr.AssemblyLines != nil {
			for _, assemblyLine := range record.AssemblyHdr.AssemblyLines {
				_, err := r.db.NewDelete().Model(assemblyLine).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
		_, err := r.db.NewDelete().Model(record.AssemblyHdr).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if record.AverageInventoryValues != nil {
		_, err := r.db.NewDelete().Model(&record.AverageInventoryValues).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
	}

	if record.InvLocMsps != nil {
		_, err := r.db.NewDelete().Model(&record.InvLocMsps).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if record.ItemIdChangeHistories != nil {
		_, err := r.db.NewDelete().Model(&record.ItemIdChangeHistories).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if record.ItemCategoryXInvMasts != nil {
		_, err := r.db.NewDelete().Model(&record.ItemCategoryXInvMasts).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if record.InventorySuppliers != nil {
		for _, invSup := range record.InventorySuppliers {
			if invSup.InventorySupplierXLocs != nil {
				_, err := r.db.NewDelete().Model(&invSup.InventorySupplierXLocs).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}
			if invSup.InventorySupplierXUoms != nil {
				_, err := r.db.NewDelete().Model(&invSup.InventorySupplierXUoms).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}

		}
		_, err := r.db.NewDelete().Model(&record.InventorySuppliers).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
	}

	if record.ItemUoms != nil {
		_, err := r.db.NewDelete().Model(&record.ItemUoms).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if record.PricePages != nil {
		for _, pricePage := range record.PricePages {
			if pricePage.PricePageXBooks != nil {
				_, err := r.db.NewDelete().Model(&pricePage.PricePageXBooks).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
		_, err := r.db.NewDelete().Model(&record.PricePages).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if record.InvXrefs != nil {
		for _, invXref := range record.InvXrefs {
			_, err := r.db.NewDelete().Model(invXref).WherePK().Exec(ctx)
			if err != nil {
				return err
			}
		}
	}

	_, err = r.db.NewDelete().Model(&record).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// applyDefaultFilters applies default filters to a query
func (r *ProductRepository) applyDefaultFilters(q *bun.SelectQuery) *bun.SelectQuery {
	if r.defaultValues.CompanyId != "" {
		q = q.Where("inv_loc.company_id = ?", r.defaultValues.CompanyId)
	}
	if r.defaultValues.LocationId != "" {
		q = q.Where("inv_loc.location_id = ?", r.defaultValues.LocationId)
	}
	return q
}

func (r *ProductRepository) selectInvMastRecord(ctx context.Context, invMastUid int) (*InvMast, error) {
	record := &InvMast{}

	err := r.db.NewSelect().Model(record).Relation("InvLocs", func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where("inv_loc.company_id = ?", r.defaultValues.CompanyId).Where("inv_loc.location_id = ?", r.defaultValues.LocationId)
	}).Where("inv_mast.inv_mast_uid = ?", invMastUid).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return record, nil
}

// safeToDeleteProduct checks if a product has any usages.
func (r *ProductRepository) safeToDeleteProduct(ctx context.Context, invMastUid int) (bool, error) {

	q := `SELECT
			(SELECT COUNT(*) FROM oe_line WHERE oe_line.inv_mast_uid = im.inv_mast_uid) as order_lines,
			(SELECT COUNT(*) FROM inv_tran WHERE inv_tran.inv_mast_uid = im.inv_mast_uid) as transactions,
			(SELECT COUNT(*) FROM inventory_receipts_line WHERE inventory_receipts_line.inv_mast_uid = im.inv_mast_uid) as receipt_lines
		FROM inv_mast im
		WHERE im.inv_mast_uid = ?`

	type result struct {
		OrderLines   int
		Transactions int
		ReceiptLines int
	}

	var res result

	err := r.db.NewRaw(q, invMastUid).Scan(ctx, &res)
	if err != nil {
		return false, err
	}

	if res.OrderLines > 0 || res.Transactions > 0 || res.ReceiptLines > 0 {
		return false, nil
	}

	return true, nil

}
