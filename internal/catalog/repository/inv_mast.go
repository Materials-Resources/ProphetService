package repository

import (
	"context"
	"fmt"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
)

type InvMast struct {
	prophet.InvMast `bun:",extend"`

	AssemblyHdr *assemblyHdr `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`

	AverageInventoryValues []*averageInventoryValue `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	AlternateCodes         []*alternateCode         `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	ItemIdChangeHistories  []*itemIdChangeHistory   `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InvAdjLines            []*invAdjLine            `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InvLocs                []*invLoc                `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InvLocMsps             []*invLocMsp             `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InvXrefs               []*invXref               `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InventorySuppliers     []*inventorySupplier     `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	ItemCategoryXInvMasts  []*itemCategoryXInvMast  `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	ItemConversions        []*itemConversion        `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	ItemUoms               []*itemUom               `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	PricePages             []*pricePage             `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
}

type invMastRepository struct {
	db bun.IDB
}

func newInvMastRepository(bun bun.IDB) invMastRepository {
	return invMastRepository{
		db: bun,
	}
}

type InvMastCreateParams struct {
	Sn          string
	Name        string
	Description string
}

func (r *invMastRepository) Select(ctx context.Context, uid int32) (*InvMast, error) {
	var record InvMast
	err := r.db.NewSelect().Model(&record).Where("inv_mast_uid = ?", uid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *invMastRepository) deleteAssociatedRecords(ctx context.Context, rec InvMast) error {

	if rec.InvLocs != nil {
		for _, invLoc := range rec.InvLocs {
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

	if rec.AlternateCodes != nil {
		for _, altCode := range rec.AlternateCodes {
			_, err := r.db.NewDelete().Model(altCode).WherePK().Exec(ctx)
			if err != nil {
				return err
			}
		}

	}

	// delete assembly_hdr records
	if rec.AssemblyHdr != nil {
		if rec.AssemblyHdr.AssemblyLines != nil {
			for _, assemblyLine := range rec.AssemblyHdr.AssemblyLines {
				_, err := r.db.NewDelete().Model(assemblyLine).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
		_, err := r.db.NewDelete().Model(rec.AssemblyHdr).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.AverageInventoryValues != nil {
		_, err := r.db.NewDelete().Model(&rec.AverageInventoryValues).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
	}

	if rec.InvLocMsps != nil {
		_, err := r.db.NewDelete().Model(&rec.InvLocMsps).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.ItemIdChangeHistories != nil {
		_, err := r.db.NewDelete().Model(&rec.ItemIdChangeHistories).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.ItemCategoryXInvMasts != nil {
		_, err := r.db.NewDelete().Model(&rec.ItemCategoryXInvMasts).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.InventorySuppliers != nil {
		for _, invSup := range rec.InventorySuppliers {
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
		_, err := r.db.NewDelete().Model(&rec.InventorySuppliers).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
	}

	if rec.ItemUoms != nil {
		_, err := r.db.NewDelete().Model(&rec.ItemUoms).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.PricePages != nil {
		for _, pricePage := range rec.PricePages {
			if pricePage.PricePageXBooks != nil {
				_, err := r.db.NewDelete().Model(&pricePage.PricePageXBooks).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
		_, err := r.db.NewDelete().Model(&rec.PricePages).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.InvXrefs != nil {
		for _, invXref := range rec.InvXrefs {
			_, err := r.db.NewDelete().Model(invXref).WherePK().Exec(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *invMastRepository) DeleteByInvMastUid(ctx context.Context, invMastUid int32) error {

	var rec InvMast

	ok, err := r.safeToDelete(ctx, invMastUid)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("product has associated records")
	}

	err = r.db.NewSelect().Model(&rec).
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

	err = r.deleteAssociatedRecords(ctx, rec)
	if err != nil {
		return err
	}
	_, err = r.db.NewDelete().Model(&rec).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *invMastRepository) safeToDelete(ctx context.Context, invMastUid int32) (bool, error) {
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
