package data

import (
	"context"
	"github.com/uptrace/bun"
)

type AlternateCodeModel struct {
	bun bun.IDB
}

func (m *AlternateCodeModel) Delete(ctx context.Context, rec *AlternateCode) error {
	_, err := m.bun.NewDelete().Model(rec).WherePK().Exec(ctx)
	return err
}

func (m *AlternateCodeModel) DeleteByInvMastUid(ctx context.Context, invMastUid int32) error {
	_, err := m.bun.NewDelete().Model((*AlternateCode)(nil)).Where("inv_mast_uid = ?", invMastUid).Exec(ctx)

	return err
}
