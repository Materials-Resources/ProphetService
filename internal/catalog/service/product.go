package service

import (
	"context"
	"fmt"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"time"
)

func NewDeleteProductWorker(sered *sr.Serde, service CatalogService) DeleteProductWorker {
	return DeleteProductWorker{sered: sered, service: service}
}

type DeleteProductWorker struct {
	sered   *sr.Serde
	service CatalogService
}

func (w DeleteProductWorker) ConsumeRecord(rec *kgo.Record) error {
	return nil
}

func NewUpdateProductWorker(sered *sr.Serde, service CatalogService) UpdateProductWorker {
	return UpdateProductWorker{sered: sered, service: service}
}

type UpdateProductWorker struct {
	sered   *sr.Serde
	service CatalogService
}

func (w UpdateProductWorker) SendToDLQ(cl *kgo.Client, rec *kgo.Record, errMsg error) error {

	var productRecord ProductRecord
	err := w.sered.Decode(rec.Value, &productRecord)
	if err != nil {
		return err
	}

	cl.Produce(
		context.Background(), &kgo.Record{
			Topic: UpdateProductDlqTopic,
			Value: w.sered.MustEncode(
				UpdateProductDlqRecord{
					Uid:   productRecord.Uid,
					Error: errMsg.Error(),
					Time:  time.Now(),
				}),
		}, func(r *kgo.Record, err error) {
			if err != nil {
				fmt.Println(err)
			}
		})
	return nil
}

func (w UpdateProductWorker) GetTopic() string {
	return UpdateProductTopic
}

func (w UpdateProductWorker) ConsumeRecord(rec *kgo.Record) error {

	var productRecord ProductRecord
	err := w.sered.Decode(rec.Value, &productRecord)
	if err != nil {
		return err
	}

	err = w.service.UpdateProduct(
		context.Background(), &domain.Product{
			UID:            productRecord.Uid,
			Name:           productRecord.Name,
			Description:    productRecord.Description,
			ProductGroupSn: productRecord.ProductGroupSn,
		}, []float64{1001})
	if err != nil {
		return err
	}
	return nil
}
