package data

import (
	"context"
	"github.com/uptrace/bun"
)

type AverageInventoryValueModel struct {
	bun bun.IDB
}

func (m *AverageInventoryValueModel) Delete(ctx context.Context, rec *AverageInventoryValue) error {
	_, err := m.bun.NewDelete().Model(rec).WherePK().Exec(ctx)
	return err
}
