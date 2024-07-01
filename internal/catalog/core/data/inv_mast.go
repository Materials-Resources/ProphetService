package data

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type InvMastModel struct {
	bun bun.IDB
}

func NewInvMastModel(bun bun.IDB) InvMastModel {
	return InvMastModel{
		bun: bun,
	}
}

type InvMastCreateParams struct {
	Sn          string
	Name        string
	Description string
}

func (m *InvMastModel) Create(ctx context.Context, params *InvMastCreateParams) error {

	return nil
}

func (m *InvMastModel) Get(ctx context.Context, invMastUid int32) (*InvMast, error) {
	var rec InvMast
	err := m.bun.NewSelect().Model(&rec).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err

	}
	return &rec, err
}
func (m *InvMastModel) deleteAssociatedRecords(ctx context.Context, rec InvMast) error {

	if rec.InvLocs != nil {
		for _, invLoc := range rec.InvLocs {
			if invLoc.InvBins != nil {
				for _, invBin := range invLoc.InvBins {
					_, err := m.bun.NewDelete().Model(invBin).WherePK().Exec(ctx)
					if err != nil {
						return err
					}
				}
			}
			if invLoc.InvLocStockStatuses != nil {
				for _, invLocStockStatus := range invLoc.InvLocStockStatuses {
					_, err := m.bun.NewDelete().Model(invLocStockStatus).WherePK().Exec(ctx)
					if err != nil {
						return err
					}
				}
			}
			_, err := m.bun.NewDelete().Model(invLoc).WherePK().Exec(ctx)
			if err != nil {
				return err
			}
		}

	}

	if rec.AlternateCodes != nil {
		for _, altCode := range rec.AlternateCodes {
			_, err := m.bun.NewDelete().Model(altCode).WherePK().Exec(ctx)
			if err != nil {
				return err
			}
		}

	}

	// delete assembly_hdr records
	if rec.AssemblyHdr != nil {
		if rec.AssemblyHdr.AssemblyLines != nil {
			for _, assemblyLine := range rec.AssemblyHdr.AssemblyLines {
				_, err := m.bun.NewDelete().Model(assemblyLine).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
		_, err := m.bun.NewDelete().Model(rec.AssemblyHdr).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.AverageInventoryValues != nil {
		_, err := m.bun.NewDelete().Model(&rec.AverageInventoryValues).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
	}

	if rec.InvLocMsps != nil {
		_, err := m.bun.NewDelete().Model(&rec.InvLocMsps).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.ItemIdChangeHistories != nil {
		_, err := m.bun.NewDelete().Model(&rec.ItemIdChangeHistories).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.ItemCategoryXInvMasts != nil {
		_, err := m.bun.NewDelete().Model(&rec.ItemCategoryXInvMasts).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.InventorySuppliers != nil {
		for _, invSup := range rec.InventorySuppliers {
			if invSup.InventorySupplierXLocs != nil {
				_, err := m.bun.NewDelete().Model(&invSup.InventorySupplierXLocs).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}
			if invSup.InventorySupplierXUoms != nil {
				_, err := m.bun.NewDelete().Model(&invSup.InventorySupplierXUoms).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}

		}
		_, err := m.bun.NewDelete().Model(&rec.InventorySuppliers).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
	}

	if rec.ItemUoms != nil {
		_, err := m.bun.NewDelete().Model(&rec.ItemUoms).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.PricePages != nil {
		for _, pricePage := range rec.PricePages {
			if pricePage.PricePageXBooks != nil {
				_, err := m.bun.NewDelete().Model(&pricePage.PricePageXBooks).WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
		_, err := m.bun.NewDelete().Model(&rec.PricePages).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	if rec.InvXrefs != nil {
		for _, invXref := range rec.InvXrefs {
			_, err := m.bun.NewDelete().Model(invXref).WherePK().Exec(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *InvMastModel) DeleteByInvMastUid(ctx context.Context, invMastUid int32) error {

	var rec InvMast

	ok, err := m.safeToDelete(ctx, invMastUid)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("product has associated records")
	}

	err = m.bun.NewSelect().Model(&rec).
		Relation("AlternateCodes").
		Relation("AssemblyHdr").
		Relation("AssemblyHdr.AssemblyLines").
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

	err = m.deleteAssociatedRecords(ctx, rec)
	if err != nil {
		return err
	}
	_, err = m.bun.NewDelete().Model(&rec).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *InvMastModel) safeToDelete(ctx context.Context, invMastUid int32) (bool, error) {
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

	err := m.bun.NewRaw(q, invMastUid).Scan(ctx, &res)
	if err != nil {
		return false, err
	}

	if res.OrderLines > 0 || res.Transactions > 0 || res.ReceiptLines > 0 {
		return false, nil
	}

	return true, nil

}
