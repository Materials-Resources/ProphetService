package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/materials-resources/s-prophet/internal/catalog/event/produce"
	"github.com/materials-resources/s-prophet/internal/catalog/repository"
	catalogv1 "github.com/materials-resources/s-prophet/internal/grpc/catalog"
	"github.com/materials-resources/s-prophet/internal/validator"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrorResourceNotFound = errors.New("resource not found")
)

func NewCatalogService(a *app.App, producer *produce.Producer) Catalog {

	return Catalog{
		app:           a,
		repository:    repository.NewRepository(a.GetDB()),
		tracer:        a.CreateTracer("service-catalog"),
		eventProducer: producer,
	}
}

type CatalogService interface {
	ListProducts(ctx context.Context, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error)
	GetProduct(ctx context.Context, req *catalogv1.GetProductRequest) (*catalogv1.GetProductResponse, error)
	// UpdateProduct produces an event to update a Product
	UpdateProduct(ctx context.Context, update *domain.ProductUpdate) error
	DeleteProduct(ctx context.Context, uid string) error

	ListGroups(ctx context.Context) ([]*domain.ProductGroup, error)
	CreateGroup(ctx context.Context, input *domain.ProductGroupInput) error
	ReadGroup(ctx context.Context, sn string) (*domain.ProductGroup, error)
	// UpdateGroup produces an event to update a productGroup
	UpdateGroup(ctx context.Context, input *domain.ProductGroupInput) error

	// GetProductBySupplierPartNumber returns a Product by the given part number and supplier
	GetProductBySupplierPartNumber(ctx context.Context, partNumber string, supplier float64) (domain.Product, error)

	UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error

	// SetPrimaryProductSupplier sets the primary product supplier
	SetPrimaryProductSupplier(ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error
}

var _ CatalogService = Catalog{}

type Catalog struct {
	repository    *repository.Repository
	app           *app.App
	tracer        trace.Tracer
	eventProducer *produce.Producer
}

func (s Catalog) UpdateProduct(ctx context.Context, product *domain.ProductUpdate) error {

	// TODO: preform authorization checks on requester

	v := validator.New()
	domain.ValidateProductUpdate(v, product)

	return s.eventProducer.UpdateProduct(ctx, product)
}

func (s Catalog) ListProducts(ctx context.Context, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error) {
	//opts := &repository.InvLocListOptions{}
	//s.repository.Product.ListProducts(ctx)
	return nil, nil, nil
}

func (s Catalog) CreateGroup(ctx context.Context, input *domain.ProductGroupInput) error {
	v := validator.New()

	input.ValidateCreate(v)

	if !v.Valid() {
		return fmt.Errorf("%s", v.Errors)
	}

	return s.repository.ProductGroup.CreateProductGroup(ctx, input)
}

func (s Catalog) ListGroups(ctx context.Context) ([]*domain.ProductGroup, error) {
	return s.repository.ProductGroup.ListProductGroups(ctx)
}

func (s Catalog) ReadGroup(ctx context.Context, sn string) (*domain.ProductGroup, error) {
	return s.repository.ProductGroup.GetProductGroup(ctx, sn)

}

func (s Catalog) HandleUpdateGroup(ctx context.Context, input *domain.ProductGroupInput) error {
	v := validator.New()

	input.ValidateUpdate(v)

	valid := v.Valid()
	if !valid {
		return fmt.Errorf("%s", v.Errors)

	}

	return s.repository.ProductGroup.UpdateProductGroup(ctx, input)
}

func (s Catalog) UpdateGroup(ctx context.Context, input *domain.ProductGroupInput) error {

	ctx, span := s.tracer.Start(ctx, "catalog.ClerkUpdateGroup")
	defer span.End()

	// TODO: preform authorization checks on requester

	v := validator.New()
	input.ValidateUpdate(v)

	return s.eventProducer.UpdateGroup(ctx, input)
}

func (s Catalog) GetProduct(ctx context.Context, req *catalogv1.GetProductRequest) (*catalogv1.GetProductResponse, error) {

	product, err := s.repository.Product.GetProduct(ctx, req.GetId())
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
			Name:        product.Name,
			Description: product.Description,
		},
	}
	return response, nil
}

func (s Catalog) SetPrimaryProductSupplier(ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error {
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

func (s Catalog) DeleteProduct(ctx context.Context, uid string) error {

	ctx, span := s.tracer.Start(ctx, "catalog.DeleteProduct")
	defer span.End()

	err := s.eventProducer.DeleteProduct(ctx, &domain.ProductUpdate{Uid: uid})
	if err != nil {
		return err
	}

	return nil
}
func (s Catalog) UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error {
	// TODO implement me
	panic("implement me")
}

func (s Catalog) GetProductBySupplierPartNumber(
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
