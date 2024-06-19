package consume

import (
	"fmt"
	"github.com/materials-resources/s-prophet/internal/catalog/core/data"
	"github.com/materials-resources/s-prophet/internal/catalog/core/event"
	"github.com/materials-resources/s-prophet/internal/catalog/core/service"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"github.com/twmb/franz-go/plugin/kotel"
	"strconv"
)

type WorkerFunc func(rec *kgo.Record) error

type Workers struct {
	tracer  *kotel.Tracer
	serde   *sr.Serde
	service service.CatalogService
	db      *data.Models
}

func NewWorkers(serde *sr.Serde, db *data.Models, tracer *kotel.Tracer, service service.CatalogService) *Workers {
	return &Workers{serde: serde, db: db, tracer: tracer, service: service}

}

func (w *Workers) UpdateGroup(rec *kgo.Record) error {
	ctx, span := w.tracer.WithProcessSpan(rec)
	defer span.End()
	var groupRecord event.ProductGroupRecord
	err := w.serde.Decode(rec.Value, &groupRecord)
	if err != nil {
		return event.ErrDecodingRecord

	}

	productGroup := event.ProductGroupDomainFromRecord(&groupRecord)

	err = w.service.UpdateGroup(ctx, &productGroup)

	return err

}

func (w *Workers) DeleteProduct(rec *kgo.Record) error {
	ctx, span := w.tracer.WithProcessSpan(rec)
	defer span.End()
	var productRecord event.ProductRecord
	err := w.serde.Decode(rec.Value, &productRecord)
	if err != nil {
		return event.ErrDecodingRecord

	}
	err = w.db.ModelsTx(ctx, func(m *data.Models) error {
		invMastUid, err := strconv.Atoi(productRecord.Uid)
		if err != nil {
			return err
		}

		err = m.InvMast.DeleteByInvMastUid(ctx, int32(invMastUid))
		if err != nil {
			return err
		}
		fmt.Println(err)
		return nil
	})
	if err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}
