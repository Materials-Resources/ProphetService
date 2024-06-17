package data

import (
	"context"
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
	// delete assembly_hdr records
	if rec.AssemblyHdr != nil {
		_, err := m.bun.NewDelete().Model(rec.AssemblyHdr.AssemblyLines).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
		_, err = m.bun.NewDelete().Model(rec.AssemblyHdr).WherePK().Exec(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}

func (m *InvMastModel) DeleteByInvMastUid(ctx context.Context, invMastUid int32) error {

	var rec InvMast
	err := m.bun.NewSelect().Model(&rec).
		Relation("AssemblyHdr").
		Relation("AssemblyHdr.AssemblyLines").
		Relation("InvLocs").
		Relation("InvLocs.InvBins").
		Relation("InvLocs.InvTrans").
		Relation("InvLocs.InvLocStockStatuses").
		Where("inv_mast.inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return err
	}

	err = m.deleteAssociatedRecords(ctx, rec)
	if err != nil {
		return err
	}
	_, err = m.bun.NewDelete().Model(rec).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
