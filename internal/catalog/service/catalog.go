package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hamba/avro/v2"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/materials-resources/s-prophet/internal/catalog/repository"
	embedavro "github.com/materials-resources/s-prophet/internal/catalog/schemas/avro"
	catalogv1 "github.com/materials-resources/s-prophet/internal/grpc/catalog"
	"github.com/materials-resources/s-prophet/pkg/helpers"
	"github.com/materials-resources/s-prophet/pkg/kafka"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"go.opentelemetry.io/otel/trace"
	"sync"
	"time"
)

var (
	ErrorResourceNotFound = errors.New("resource not found")
)

var (
	TopicUpdateProduct = "PROPHET_CATALOG_UPDATE_PRODUCT"
	TopicUpdateGroup   = "PROPHET_CATALOG_UPDATE_GROUP"
	TopicDeleteProduct = "PROPHET_CATALOG_DELETE_PRODUCT"
	TopicDeleteGroup   = "PROPHET_CATALOG_DELETE_GROUP"
	TopicCreateGroup   = "PROPHET_CATALOG_CREATE_GROUP"
	TopicCreateProduct = "PROPHET_CATALOG_CREATE_PRODUCT"
)

func NewCatalogService(a *app.App, kafkaFactory *kafka.Factory, srClient *sr.Client) Catalog {

	clientID := "PROPHET_CATALOG_SERVICE"

	validate := validator.New(validator.WithRequiredStructEnabled())

	kProducer, err := kafkaFactory.NewKafkaClient(clientID)
	if err != nil {
		panic(err)
	}
	kConsumer, err := kafkaFactory.NewKafkaClient(clientID, kgo.ConsumerGroup(clientID),
		kgo.ConsumeTopics("PROPHET_CATALOG_UPDATE_PRODUCT"))
	if err != nil {
		panic(err)
	}

	return Catalog{
		app:           a,
		repository:    repository.NewRepository(a.GetDB()),
		tracer:        a.Otel.GetTracerProvider().Tracer("service-catalog"),
		validate:      validate,
		kProducer:     kProducer,
		kConsumer:     kConsumer,
		kFactory:      kafkaFactory,
		consumerGroup: kafka.NewConsumerGroup(),
		srClient:      srClient,
		avroSerde:     &sr.Serde{},
	}
}

type CatalogService interface {
	ListProducts(ctx context.Context, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error)
	GetProduct(ctx context.Context, req *catalogv1.GetProductRequest) (*catalogv1.GetProductResponse, error)
	// UpdateProduct produces an event to update a Product
	UpdateProduct(ctx context.Context, req *catalogv1.UpdateProductRequest) (*catalogv1.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, uid string) error

	ListGroups(ctx context.Context) ([]*domain.ProductGroup, error)
	CreateGroup(ctx context.Context, req *catalogv1.CreateGroupRequest) (*catalogv1.CreateGroupResponse, error)
	ReadGroup(ctx context.Context, sn string) (*domain.ProductGroup, error)
	// UpdateGroup produces an event to update a productGroup
	UpdateGroup(ctx context.Context, req *catalogv1.UpdateGroupRequest) (*catalogv1.UpdateGroupResponse, error)

	// GetProductBySupplierPartNumber returns a Product by the given part number and supplier
	GetProductBySupplierPartNumber(ctx context.Context, partNumber string, supplier float64) (domain.Product, error)

	UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error

	// SetPrimaryProductSupplier sets the primary product supplier
	SetPrimaryProductSupplier(ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error
}

var _ CatalogService = &Catalog{}

type Catalog struct {
	repository *repository.Repository
	app        *app.App
	tracer     trace.Tracer
	validate   *validator.Validate

	kConsumer     *kgo.Client
	kFactory      *kafka.Factory
	consumerGroup kafka.ConsumerGroup
	kProducer     *kgo.Client

	avroSerde *sr.Serde
	srClient  *sr.Client
}

func (svc *Catalog) Initialize(ctx context.Context) error {

	err := svc.reconcileTopic(ctx, TopicUpdateProduct)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}
	err = svc.reconcileTopic(ctx, TopicCreateProduct)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}
	err = svc.reconcileTopic(ctx, TopicDeleteProduct)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}
	err = svc.reconcileTopic(ctx, TopicCreateGroup)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}
	err = svc.reconcileTopic(ctx, TopicUpdateGroup)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}
	err = svc.reconcileTopic(ctx, TopicDeleteGroup)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}

	if svc.srClient != nil {
		productSchemaID, err := svc.registerAvroSchema(ctx, "org.materialsresources.v1.catalog.avro.Product",
			embedavro.ProductAvro)
		if err != nil {
			return fmt.Errorf("failed to register product schema: %w", err)
		}

		groupSchemaID, err := svc.registerAvroSchema(ctx, "org.materialsresources.v1.catalog.avro.Group", embedavro.GroupAvro)
		if err != nil {
			return fmt.Errorf("failed to register group schema: %w", err)
		}

		// Parse all schemas to add them to the global cache
		avro.DefaultConfig = avro.Config{
			TagKey: "avro",
		}.Freeze()
		productAvroSchema, err := avro.Parse(embedavro.ProductAvro)
		if err != nil {
			return fmt.Errorf("failed to parse product avro schema with avro lib: %w", err)
		}

		groupAvroSchema, err := avro.Parse(embedavro.GroupAvro)
		if err != nil {
			return fmt.Errorf("failed to parse group avro schema with avro lib: %w", err)
		}

		svc.avroSerde.Register(productSchemaID, domain.Product{}, sr.EncodeFn(func(v any) ([]byte, error) {
			return avro.Marshal(productAvroSchema, v)
		}), sr.DecodeFn(
			func(bytes []byte, a any) error {
				return avro.Unmarshal(productAvroSchema, bytes, a)
			}))

		svc.avroSerde.Register(groupSchemaID, domain.ProductGroup{}, sr.EncodeFn(func(v any) ([]byte, error) {
			return avro.Marshal(groupAvroSchema, v)
		}), sr.DecodeFn(
			func(bytes []byte, a any) error {
				return avro.Unmarshal(groupAvroSchema, bytes, a)
			}))
	}

	svc.consumerGroup.AddWorker(TopicUpdateGroup, svc.consumeUpdateGroup)

	err = svc.consumerGroup.Consume("PROPHET_CATALOG_SERVICE", "PROPHET_CATALOG_SERVICE_1", *svc.kFactory)
	if err != nil {
		return fmt.Errorf("failed to start consumer group from kafka: %w", err)
	}

	return nil

}

func (svc *Catalog) UpdateProduct(ctx context.Context, req *catalogv1.UpdateProductRequest) (*catalogv1.
	UpdateProductResponse, error) {

	product := &domain.Product{
		Create:      false,
		Sn:          req.GetProduct().Sn,
		Name:        &req.GetProduct().Name,
		Description: &req.GetProduct().Description,
	}

	err := svc.validate.Struct(product)
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)
	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("%s", validationErrors)
	}

	err = svc.produceUpdateProduct(product)

	return &catalogv1.UpdateProductResponse{}, nil
}

func (svc *Catalog) produceUpdateProduct(product *domain.Product) error {
	return nil

}
func (svc *Catalog) ListProducts(ctx context.Context, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error) {
	//opts := &repository.InvLocListOptions{}
	//s.repository.Product.ListProducts(ctx)
	return nil, nil, nil
}

func (svc *Catalog) CreateGroup(ctx context.Context, req *catalogv1.CreateGroupRequest) (*catalogv1.
	CreateGroupResponse, error) {

	productGroup := &domain.ProductGroup{
		Create: true,
		Sn:     req.GetSn(),
		Name:   &req.Name,
	}

	err := svc.validate.Struct(productGroup)
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)

	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("%s", validationErrors)
	}

	err = svc.repository.ProductGroup.CreateProductGroup(ctx, productGroup)
	if err != nil {
		return nil, err
	}

	return &catalogv1.CreateGroupResponse{}, nil
}

func (svc *Catalog) ListGroups(ctx context.Context) ([]*domain.ProductGroup, error) {
	return svc.repository.ProductGroup.ListProductGroups(ctx)
}

func (svc *Catalog) ReadGroup(ctx context.Context, sn string) (*domain.ProductGroup, error) {
	return svc.repository.ProductGroup.GetProductGroup(ctx, sn)

}

func (svc *Catalog) UpdateGroup(ctx context.Context, req *catalogv1.UpdateGroupRequest) (*catalogv1.
	UpdateGroupResponse, error) {

	productGroup := &domain.ProductGroup{
		Create: false,
		Sn:     req.Sn,
	}

	if req.GetProductGroup().GetUpdateMask() != nil {
		for _, path := range req.GetProductGroup().UpdateMask.Paths {
			switch path {
			case "name":
				productGroup.Name = &req.GetProductGroup().Name
			}
		}
	}

	err := svc.validate.Struct(productGroup)
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)

	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("%s", validationErrors)
	}

	err = svc.produceUpdateGroup(ctx, *productGroup)

	if err != nil {
		return nil, err
	}

	return &catalogv1.UpdateGroupResponse{}, nil
}

func (svc *Catalog) produceUpdateGroup(ctx context.Context, productGroup domain.ProductGroup) error {

	span := trace.SpanFromContext(ctx)

	value, err := svc.avroSerde.Encode(productGroup)
	if err != nil {
		span.RecordError(fmt.Errorf("failed to encode product group: %w", err)) // Ensure error is recorded in tracing
		return err
	}

	rec := &kgo.Record{
		Topic:     TopicUpdateGroup,
		Key:       []byte(productGroup.Sn),
		Value:     value,
		Timestamp: time.Now(),
	}

	var wg sync.WaitGroup
	wg.Add(1)

	svc.kProducer.Produce(ctx, rec, func(record *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			span.RecordError(fmt.Errorf("record had a produce error: %v\n", err))
			return
		}
	})
	wg.Wait()

	return nil
}

func (svc *Catalog) consumeUpdateGroup(rec *kgo.Record) error {
	ctx, _ := svc.kFactory.GetKotelTracer().WithProcessSpan(rec)

	group := &domain.ProductGroup{}
	err := svc.avroSerde.Decode(rec.Value, group)
	if err != nil {
		return fmt.Errorf("failed to decode record: %w", err)
	}
	err = svc.repository.ProductGroup.UpdateProductGroup(ctx, group)
	if err != nil {
		return err
	}
	return nil

}
func (svc *Catalog) GetProduct(ctx context.Context, req *catalogv1.GetProductRequest) (*catalogv1.GetProductResponse, error) {

	product, err := svc.repository.Product.GetProduct(ctx, req.GetId())
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorResourceNotFound
		}
		return nil, err
	}

	response := &catalogv1.GetProductResponse{
		Product: &catalogv1.Product{
			Id:          product.Id,
			Sn:          product.Sn,
			Name:        helpers.GetOptionalValue(product.Name, ""),
			Description: helpers.GetOptionalValue(product.Description, ""),
		},
	}
	return response, nil
}

func (svc *Catalog) SetPrimaryProductSupplier(ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error {
	//ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error {
	//
	//err := s.m.ExecTx(
	//	ctx, func(m *models.Models) error {
	//
	//		inventorySupplier, err := m.InventorySupplier.GetBySupplierIdDivisionIdInvMastUid(
	//			ctx, supplierUid,
	//			divisionId, productUid)
	//		if err != nil {
	//			return err
	//
	//		}
	//		inventorySupplierXLocs, err := m.InventorySupplierXLoc.GetByInvMastUidAndLocationId(
	//			ctx, productUid,
	//			locationId)
	//
	//		if err != nil {
	//			return err
	//		}
	//
	//		for _, inventorySupplierXLoc := range inventorySupplierXLocs {
	//			if inventorySupplierXLoc.InventorySupplierUid == inventorySupplier.InventorySupplierUid {
	//				inventorySupplierXLoc.PrimarySupplier = "Y"
	//			} else {
	//				inventorySupplierXLoc.PrimarySupplier = "N"
	//			}
	//			if err := m.InventorySupplierXLoc.Update(ctx, inventorySupplierXLoc); err != nil {
	//				return err
	//			}
	//
	//		}
	//		return nil
	//	})
	//
	//return err

	return nil

}

func (svc *Catalog) DeleteProduct(ctx context.Context, uid string) error {
	//
	//err := svc.eventProducer.DeleteProduct(ctx, &domain.ProductUpdate{Uid: uid})
	//if err != nil {
	//	return err
	//}

	return nil
}
func (svc *Catalog) UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error {
	// TODO implement me
	panic("implement me")
}

func (svc *Catalog) GetProductBySupplierPartNumber(
	ctx context.Context, partNumber string,
	supplier float64) (domain.Product, error) {
	//invMast, err := s.repository.Product..GetBySupplierPartNumber(ctx, supplier, partNumber)
	//if err != nil {
	//	return domain.Product{}, err
	//}
	//return domain.Product{
	//	Id: strconv.Itoa(int(invMast.InvMastUid)), Sn: &invMast.ItemId, Name: &invMast.ItemDesc,
	//}, nil
	// TODO implement me
	panic("implement me")
}

func (svc *Catalog) reconcileTopic(ctx context.Context, topic string) error {
	topicCfg := map[string]*string{
		"cleanup.policy": kadm.StringPtr("compact"),
	}
	err := kafka.ReconcileTopic(ctx,
		svc.kProducer,
		topic,
		4,
		1,
		topicCfg,
	)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}
	return nil
}

func (svc *Catalog) registerAvroSchema(ctx context.Context, subject, schema string) (int, error) {
	registeredSchema, err := svc.srClient.CreateSchema(
		ctx, subject, sr.Schema{
			Schema: schema,
			Type:   sr.TypeAvro,
		})

	if err != nil {
		return -1, err
	}
	return registeredSchema.ID, nil

}
