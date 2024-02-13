package kafka

import (
	"context"

	"github.com/materials-resources/s_prophet/internal/catalog"
	"github.com/twmb/franz-go/pkg/kgo"
)

type ProductWorker struct {
	repo catalog.Repository
}

func NewProductWorker(repo catalog.Repository) *ProductWorker {
	return &ProductWorker{repo: repo}
}

func (w *ProductWorker) ConsumeDeleteProduct(ctx context.Context, rec *kgo.Record) error {
	return w.repo.DeleteProduct(ctx, string(rec.Value))
}
