package service

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"time"
)

type EventProducer struct {
	client *kgo.Client
	serde  *sr.Serde
}

type KafkaProducer struct {
	Client *kgo.Client
	Serde  *sr.Serde
}

type ProductRecord struct {
	Uid            string `avro:"uid"`
	Name           string `avro:"name"`
	Description    string `avro:"description"`
	ProductGroupSn string `avro:"product_group_sn"`
}

type UpdateProductDlqRecord struct {
	Uid   string    `avro:"uid"`
	Error string    `avro:"error"`
	Time  time.Time `avro:"time"`
}

func (p *EventProducer) PublishDeleteProduct(ctx context.Context, invMastUid string) error {

	p.client.Produce(
		context.Background(), &kgo.Record{
			Topic: DeleteProductTopic,
			Value: p.serde.MustEncode(
				ProductRecord{
					Uid: invMastUid,
				}),
		}, func(record *kgo.Record, err error) {
			if err != nil {
				fmt.Printf("record had a produce error: %v\n", err)
			}

		})
	return nil
}
func (p *KafkaProducer) PublishUpdateProduct(ctx context.Context, product *domain.Product) error {

	p.Client.Produce(
		context.Background(), &kgo.Record{
			Topic: UpdateProductTopic,
			Value: p.Serde.MustEncode(
				ProductRecord{
					Uid:            product.Uid,
					Name:           *product.Name,
					Description:    *product.Description,
					ProductGroupSn: *product.ProductGroupSn,
				}),
		}, func(record *kgo.Record, err error) {
			if err != nil {
				fmt.Printf("record had a produce error: %v\n", err)
			}

		})
	return nil
}
