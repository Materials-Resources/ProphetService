package service

import (
	"context"
	"fmt"
	"github.com/hamba/avro/v2"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"time"
)

const (
	UpdateProductTopic    = "product_update"
	UpdateProductDlqTopic = "product_update_dlq"
)

const (
	productSchemaText = `{
	"type": "record",
	"name": "product_update",
	"namespace": "org.hamba.avro",
	"fields" : [
		{"name": "uid", "type": "int"},
		{"name": "name", "type": "string"},
		{"name": "description", "type": "string"},
		{"name": "product_group_sn", "type": "string"}
	]
}`
	updateProductDlqSchema = `{
	"type": "record",
	"name": "product_update_dlq",
	"namespace": "org.hamba.avro",
	"fields" : [
		{"name": "uid", "type": "int"},
		{"name": "error", "type": "string"},
		{"name": "time", "type": "string", "logicalType": "timestamp-millis"}
	]
}
`
)

type KafkaProducer struct {
	Client *kgo.Client
	Serde  *sr.Serde
}

func RegisterSchema(serde *sr.Serde) error {
	rcl, err := sr.NewClient(sr.URLs("localhost:18081"))
	if err != nil {
		return err
	}
	register("update_product", productSchemaText, ProductRecord{}, serde, err, rcl)
	register("update_product_dlq", updateProductDlqSchema, UpdateProductDlqRecord{}, serde, err, rcl)

	return nil
}

func register(subject, schema string, schemaStruct any, serde *sr.Serde, err error, rcl *sr.Client) {
	ss, err := rcl.CreateSchema(
		context.Background(), subject, sr.Schema{
			Schema: schema,
			Type:   sr.TypeAvro,
		})
	if err != nil {
		fmt.Println(err)
	}
	avroSchema, err := avro.Parse(schema)
	if err != nil {
		fmt.Println(err)
	}
	serde.Register(
		ss.ID, schemaStruct, sr.EncodeFn(
			func(v any) ([]byte, error) {
				return avro.Marshal(avroSchema, v)
			}), sr.DecodeFn(
			func(bytes []byte, a any) error {
				return avro.Unmarshal(avroSchema, bytes, a)
			}),
	)
}

type ProductRecord struct {
	Uid            int32  `avro:"uid"`
	Name           string `avro:"name"`
	Description    string `avro:"description"`
	ProductGroupSn string `avro:"product_group_sn"`
}

type UpdateProductDlqRecord struct {
	Uid   int32     `avro:"uid"`
	Error string    `avro:"error"`
	Time  time.Time `avro:"time"`
}

func (p *KafkaProducer) PublishUpdateProduct(ctx context.Context, product *domain.Product) error {

	p.Client.Produce(
		context.Background(), &kgo.Record{
			Topic: UpdateProductTopic,
			Value: p.Serde.MustEncode(
				ProductRecord{
					Uid:            product.UID,
					Name:           product.Name,
					Description:    product.Description,
					ProductGroupSn: product.ProductGroupSn,
				}),
		}, func(record *kgo.Record, err error) {
			if err != nil {
				fmt.Printf("record had a produce error: %v\n", err)
			}

		})
	return nil
}
