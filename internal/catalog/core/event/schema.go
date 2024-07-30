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

type ProductRecord struct {
	Uid            string  `avro:"uid"`
	Name           *string `avro:"name"`
	Description    *string `avro:"description"`
	ProductGroupSn *string `avro:"product_group_sn"`
}

type ProductGroupRecord struct {
	Sn   string  `avro:"sn"`
	Name *string `avro:"name"`
}

var schemas = []struct {
	name         string
	schema       string
	schemaStruct any
}{
	{
		name: "product",
		schema: `{
	"type": "record",
	"name": "product",
	"namespace": "org.hamba.avro",
	"fields" : [
		{"name": "uid", "type": "string"},
		{"name": "name", "type": ["string", "null"]},
		{"name": "description", "type": ["string", "null"]},
		{"name": "product_group_sn", "type": ["string", "null"]}
	]
}`,
		schemaStruct: ProductRecord{},
	},
	{
		name: "product_group",
		schema: `{
	"type": "record",
	"name": "product_group",
	"namespace": "org.materialsresources.avro",
	"fields" : [
		{"name": "sn", "type": "string"},
		{"name": "name", "type": ["string", "null"]}
	]
}`,
		schemaStruct: ProductGroupRecord{},
	},
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

// RegisterSchemasWithRegistry registers schemas with the kafka schema registry
func (m *Manager) RegisterSchemasWithRegistry() error {
	cl, err := sr.NewClient(sr.URLs(m.app.Config().Kafka.Registry...))
	if err != nil {
		return fmt.Errorf("could not create schema registry client: %w", err)
	}

	for _, s := range schemas {
		ctx := context.Background()
		_, err := cl.CreateSchema(ctx, s.name, sr.Schema{
			Type:   sr.TypeAvro,
			Schema: s.schema,
		})
		if err != nil {
			return fmt.Errorf("could not create schema: %w", err)
		}
	}

	return nil

}

// RegisterSchemasWithSerde associates schemas with serde
func (m *Manager) RegisterSchemasWithSerde() error {

	// Create a new connection to the schema registry
	cl, err := sr.NewClient(sr.URLs(m.app.Config().Kafka.Registry...))
	if err != nil {
		return fmt.Errorf("could not create schema registry client: %w", err)
	}

	for _, s := range schemas {
		ctx := context.Background()

		// Lookup the schema in the registry
		ss, err := cl.LookupSchema(ctx, s.name, sr.Schema{
			Type:   sr.TypeAvro,
			Schema: s.schema,
		})

		avroSchema, err := avro.Parse(s.schema)
		if err != nil {
			return fmt.Errorf("could not parse schema: %w", err)
		}

		// Register the schema with the serde
		m.Serde.Register(
			ss.ID, s.schemaStruct, sr.EncodeFn(
				func(v any) ([]byte, error) {
					return avro.Marshal(avroSchema, v)
				}), sr.DecodeFn(
				func(bytes []byte, a any) error {
					return avro.Unmarshal(avroSchema, bytes, a)
				}),
		)

	}
	return nil

}
