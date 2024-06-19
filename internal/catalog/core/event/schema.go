package event

import (
	"context"
	"errors"
	"fmt"
	"github.com/hamba/avro/v2"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/twmb/franz-go/pkg/sr"
)

var (
	ErrDecodingRecord = errors.New("could not decode record")
)

const (
	productSchema = `{
	"type": "record",
	"name": "product",
	"namespace": "org.hamba.avro",
	"fields" : [
		{"name": "uid", "type": "string"},
		{"name": "name", "type": ["string", "null"]},
		{"name": "description", "type": ["string", "null"]},
		{"name": "product_group_sn", "type": ["string", "null"]}
	]
}`
)

type ProductRecord struct {
	Uid            string  `avro:"uid"`
	Name           *string `avro:"name"`
	Description    *string `avro:"description"`
	ProductGroupSn *string `avro:"product_group_sn"`
}

func ProductRecordFromDomain(product *domain.Product) ProductRecord {
	return ProductRecord{
		Uid:            product.Uid,
		Name:           product.Name,
		Description:    product.Description,
		ProductGroupSn: product.ProductGroupSn,
	}
}

func ProductDomainFromRecord(record ProductRecord) domain.Product {
	return domain.Product{
		Uid:            record.Uid,
		Name:           record.Name,
		Description:    record.Description,
		ProductGroupSn: record.ProductGroupSn,
	}
}

const productGroupSchema = `{
	"type": "record",
	"name": "product_group",
	"namespace": "org.materialsresources.avro",
	"fields" : [
		{"name": "sn", "type": "string"},
		{"name": "name", "type": ["string", "null"]}
	]
}`

type ProductGroupRecord struct {
	Sn   string  `avro:"sn"`
	Name *string `avro:"name"`
}

func ProductGroupRecordFromDomain(productGroup *domain.ProductGroup) ProductGroupRecord {
	return ProductGroupRecord{
		Sn:   productGroup.Sn,
		Name: productGroup.Name,
	}
}

func ProductGroupDomainFromRecord(record *ProductGroupRecord) domain.ProductGroup {
	return domain.ProductGroup{
		Sn:   record.Sn,
		Name: record.Name,
	}

}

type Registry struct {
	client *sr.Client
}

func NewRegistry() *Registry {
	cl, err := sr.NewClient(sr.URLs("localhost:18081"))
	if err != nil {
		fmt.Println(err)
	}
	return &Registry{client: cl}
}

func (r *Registry) RegisterSchemas(serde *sr.Serde) error {
	err := r.registerSchema("product", productSchema, ProductRecord{}, serde)
	if err != nil {
		return err
	}

	err = r.registerSchema("product_group", productGroupSchema, ProductGroupRecord{}, serde)
	if err != nil {
		return err
	}
	return nil

}
func (r *Registry) registerSchema(name, schema string, schemaStruct any, serde *sr.Serde) error {
	ss, err := r.client.CreateSchema(
		context.Background(), name, sr.Schema{
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
	return nil
}
