package data

import (
	"context"
	"github.com/uptrace/bun"
)

type InventoryReceiptsLineModel struct {
	bun bun.IDB
}

func (m *InventoryReceiptsLineModel) Delete(ctx context.Context, rec *InventoryReceiptsLine) error {
	_, err := m.bun.NewDelete().Model(rec).WherePK().Exec(ctx)
	return err
}
