package service

import (
	"context"
	"fmt"
	"github.com/hamba/avro/v2"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
)

const productSchemaText = `{
	"type": "record",
	"name": "simple",
	"namespace": "org.hamba.avro",
	"fields" : [
		{"name": "uid", "type": "int"},
		{"name": "name", "type": "string"},
		{"name": "description", "type": "string"},
		{"name": "product_group_sn", "type": "string"}
	]
}`

type KafkaProducer struct {
	client *kgo.Client
	serde  sr.Serde
}

func RegisterSchema(serde *sr.Serde) error {
	rcl, err := sr.NewClient(sr.URLs("localhost:18081"))
	if err != nil {
		return err
	}
	ss, err := rcl.CreateSchema(
		context.Background(), "product", sr.Schema{
			Schema: productSchemaText,
			Type:   sr.TypeAvro,
		})
	if err != nil {
		fmt.Println(err)
	}
	avroSchema, err := avro.Parse(productSchemaText)
	if err != nil {
		fmt.Println(err)
	}
	serde.Register(
		ss.ID, UpdateProductRecord{}, sr.EncodeFn(
			func(v any) ([]byte, error) {
				return avro.Marshal(avroSchema, v)
			}), sr.DecodeFn(
			func(bytes []byte, a any) error {
				return avro.Unmarshal(avroSchema, bytes, a)
			}),
	)
	return nil
}

type UpdateProductRecord struct {
	Uid            int32  `avro:"uid"`
	Name           string `avro:"name"`
	Description    string `avro:"description"`
	ProductGroupSn string `avro:"product_group_sn"`
}

func (p *KafkaProducer) PublishUpdateProduct(ctx context.Context, product *domain.Product) error {

	// rec := &UpdateProductRecord{
	// 	Uid:            product.UID,
	// 	Name:           product.Name,
	// 	Description:    product.Description,
	// 	ProductGroupSn: product.ProductGroupSn,
	// }

	// p.client.Produce(ctx, &kgo.Record{})
	return nil
}
