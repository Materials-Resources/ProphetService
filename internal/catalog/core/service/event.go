package service

import (
	"context"
	"fmt"
	"github.com/hamba/avro/v2"
	"github.com/twmb/franz-go/pkg/sr"
)

const (
	UpdateProductTopic    = "product_update"
	UpdateProductDlqTopic = "product_update_dlq"
	DeleteProductTopic    = "product_delete"
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
