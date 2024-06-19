package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/core/data"
	"github.com/materials-resources/s-prophet/internal/catalog/core/event/produce"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/materials-resources/s-prophet/internal/validator"
	"github.com/materials-resources/s-prophet/pkg/models"
	"go.opentelemetry.io/otel/trace"
	"strconv"
)

var (
	ErrorResourceNotFound = errors.New("resource not found")
)

func NewCatalogService(a *app.App, producer *produce.Producer) Catalog {

	return Catalog{
		app:           a,
		localModel:    data.NewModels(a.GetDB()),
		m:             a.GetModels(),
		tracer:        a.CreateTracer("service-catalog"),
		eventProducer: producer,
	}
}

type Catalog struct {
	app           *app.App
	m             models.Models
	tracer        trace.Tracer
	localModel    *data.Models
	eventProducer *produce.Producer
}

func (c Catalog) UpdateGroup(ctx context.Context, productGroup *domain.ProductGroup) error {
	v := validator.New()
	domain.ValidateProductGroupUpdate(v, *productGroup)

	valid := v.Valid()
	if !valid {
		return errors.New(fmt.Sprintf("%s", v.Errors))

	}

	return c.localModel.ProductGroup.Update(ctx, productGroup)
}

func (c Catalog) UpdateProduct(ctx context.Context, product *domain.Product, locations []float64) error {
	//TODO implement me
	panic("implement me")
}

func (c Catalog) ClerkUpdateProduct(ctx context.Context, product *domain.Product) error {
	v := validator.New()

	domain.ValidateProductUpdate(v, *product)

	if !v.Valid() {
		return errors.New(fmt.Sprintf("%s", v.Errors))

	}

	return c.localModel.InvLoc.Update(ctx, product)
}

func (c Catalog) ClerkDeleteProduct(ctx context.Context, uid string, deleteMode domain.DeleteMode) error {
	//TODO implement me
	panic("implement me")
}

func (c Catalog) ClerkListProducts(ctx context.Context, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error) {
	opts := &data.InvLocListOptions{}
	return c.localModel.InvLoc.List(ctx, opts, pagination)
}

func (c Catalog) ClerkCreateGroup(ctx context.Context, productGroup *domain.ProductGroup) error {
	v := validator.New()

	domain.ValidateProductGroupCreate(v, *productGroup)

	if !v.Valid() {
		return errors.New(fmt.Sprintf("%s", v.Errors))
	}

	return c.localModel.ProductGroup.Create(ctx, productGroup)
}

func (c Catalog) ClerkListGroups(ctx context.Context) ([]*domain.ProductGroup, error) {
	opts := &data.ProductGroupListOptions{}
	return c.localModel.ProductGroup.List(ctx, opts)
}

func (c Catalog) ClerkReadGroup(ctx context.Context, sn string) (*domain.ProductGroup, error) {
	return c.localModel.ProductGroup.Get(ctx, sn)

}

func (c Catalog) ClerkUpdateGroup(ctx context.Context, productGroup *domain.ProductGroup) error {

	ctx, span := c.tracer.Start(ctx, "catalog.ClerkUpdateGroup")
	defer span.End()

	// TODO: preform authorization checks on requester

	v := validator.New()
	domain.ValidateProductGroupUpdate(v, *productGroup)

	return c.eventProducer.UpdateGroup(ctx, productGroup)
}

func (c Catalog) ClerkGetProduct(ctx context.Context, uid string) (*domain.Product, error) {
	return c.getProduct(ctx, uid)
}

func (c Catalog) CustomerGetProduct(ctx context.Context, uid string) (*domain.Product, error) {
	return c.getProduct(ctx, uid)
}

func (c Catalog) getProduct(ctx context.Context, uid string) (*domain.Product, error) {

	product, err := c.localModel.InvLoc.Get(ctx, uid)
	if err != nil {
		return nil, err

	}
	return product, nil
}

func (c Catalog) SetPrimaryProductSupplier(
	ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error {

	err := c.m.ExecTx(
		ctx, func(m *models.Models) error {

			inventorySupplier, err := m.InventorySupplier.GetBySupplierIdDivisionIdInvMastUid(
				ctx, supplierUid,
				divisionId, productUid)
			if err != nil {
				return err

			}
			inventorySupplierXLocs, err := m.InventorySupplierXLoc.GetByInvMastUidAndLocationId(
				ctx, productUid,
				locationId)

			for _, inventorySupplierXLoc := range inventorySupplierXLocs {
				if inventorySupplierXLoc.InventorySupplierUid == inventorySupplier.InventorySupplierUid {
					inventorySupplierXLoc.PrimarySupplier = "Y"
				} else {
					inventorySupplierXLoc.PrimarySupplier = "N"
				}
				fmt.Println(
					inventorySupplier.InventorySupplierUid, inventorySupplierXLoc.InventorySupplierUid,
					inventorySupplierXLoc.PrimarySupplier)
				if err := m.InventorySupplierXLoc.Update(ctx, inventorySupplierXLoc); err != nil {
					return err
				}

			}
			return nil
		})

	return err
}

func (c Catalog) RegisterWorkers() {

	//cg := consumer.NewConsumerGroup()
	//
	//cg.RegisterWorkers(
	//	consumer.NewWorker(
	//		DeleteProductTopic, func(rec *kgo.Record) error {
	//			return c.event.Consumer.ConsumeDeleteProduct(context.Background(), rec, c.DeleteProduct)
	//
	//		}))
	//
	//cg.Start([]string{"localhost:19092"}, "18")
}

func (c Catalog) DeleteProduct(ctx context.Context, uid string) error {

	ctx, span := c.tracer.Start(ctx, "catalog.DeleteProduct")
	defer span.End()

	err := c.eventProducer.DeleteProduct(ctx, &domain.Product{Uid: uid})
	if err != nil {
		return err
	}

	return nil
}
func (c Catalog) UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error {
	// TODO implement me
	panic("implement me")
}

func (c Catalog) GetProductBySupplierPartNumber(
	ctx context.Context, partNumber string,
	supplier float64) (domain.Product, error) {
	invMast, err := c.m.InvMast.GetBySupplierPartNumber(ctx, supplier, partNumber)
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		Uid: strconv.Itoa(int(invMast.InvMastUid)), Sn: &invMast.ItemId, Name: &invMast.ItemDesc,
	}, nil
}
