package consume

import (
	"github.com/materials-resources/s-prophet/internal/catalog/event"
	"github.com/materials-resources/s-prophet/internal/catalog/repository"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"github.com/twmb/franz-go/plugin/kotel"
	"go.opentelemetry.io/otel/codes"
)

type WorkerFunc func(rec *kgo.Record) error

type Workers struct {
	tracer     *kotel.Tracer
	serde      *sr.Serde
	repository *repository.Repository
}

func NewWorkers(serde *sr.Serde, tracer *kotel.Tracer, repo *repository.Repository) *Workers {
	return &Workers{serde: serde, tracer: tracer, repository: repo}

}

func (w *Workers) UpdateGroup(rec *kgo.Record) error {
	ctx, span := w.tracer.WithProcessSpan(rec)
	defer span.End()
	var groupRecord event.ProductGroupInputRecord
	err := w.serde.Decode(rec.Value, &groupRecord)
	if err != nil {
		return event.ErrDecodingRecord

	}

	err = w.repository.ProductGroup.UpdateProductGroup(ctx, event.ProductGroupDomainFromRecord(&groupRecord))

	return err

}

func (w *Workers) DeleteProduct(rec *kgo.Record) error {
	ctx, span := w.tracer.WithProcessSpan(rec)
	defer span.End()
	var productUpdateRecord event.ProductUpdateRecord
	err := w.serde.Decode(rec.Value, &productUpdateRecord)
	if err != nil {
		return event.ErrDecodingRecord

	}

	err = w.repository.Product.DeleteProduct(ctx, event.ProductUpdateFromRecord(productUpdateRecord).Uid)

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	return nil
}

func (w *Workers) UpdateProduct(rec *kgo.Record) error {
	ctx, span := w.tracer.WithProcessSpan(rec)
	defer span.End()
	var productUpdateRecord event.ProductUpdateRecord
	err := w.serde.Decode(rec.Value, &productUpdateRecord)
	if err != nil {
		return event.ErrDecodingRecord

	}

	err = w.repository.Product.UpdateProduct(ctx, event.ProductUpdateFromRecord(productUpdateRecord))

	return err
}
