package data

import (
	"context"
	"github.com/uptrace/bun"
)

type AssemblyHdrModel struct {
	bun bun.IDB
}

func (m *AssemblyHdrModel) Delete(ctx context.Context, rec *AssemblyHdr) error {
	_, err := m.bun.NewDelete().Model(rec).WherePK().Exec(ctx)
	return err
}

func (m *AssemblyHdrModel) DeleteByInvMastUid(ctx context.Context, invMastUid int32) error {
	_, err := m.bun.NewDelete().Model((*AssemblyHdr)(nil)).Where("inv_mast_uid = ?", invMastUid).Exec(ctx)
	return err
}
