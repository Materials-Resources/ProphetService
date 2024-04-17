package prophet_21_1_4559

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type Queries struct {
	db bun.DB
}

func (m *Models) ExecTx(ctx context.Context, fn func(*Models) error) error {
	tx, err := m.Bun.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := NewModels(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
